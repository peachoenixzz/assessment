package api

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peachoenixz/assessment/internal/database"
	"github.com/peachoenixz/assessment/internal/expense"
	"github.com/peachoenixz/assessment/pkg/log"
)

type RouterUseCase interface {
	routerRead(endpoint *expense.Endpoint)
	routerUpdate(endpoint *expense.Endpoint)
	routerCreate(endpoint *expense.Endpoint)
}

type RouterSession struct {
	Session *echo.Echo
	Client  *sql.DB
}

func Auth() echo.MiddlewareFunc {
	res := middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("apidesign")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("123456")) == 1 {
			return true, nil
		}
		return false, nil
	})
	return res
}

func (r RouterSession) routerRead(endpoint *expense.Endpoint) {
	r.Session.GET("/expenses/:id", endpoint.ViewExpenseByID)
	r.Session.GET("/expenses", endpoint.ViewExpense, Auth())
}

func (r RouterSession) routerUpdate(endpoint *expense.Endpoint) {
	r.Session.PUT("/expenses/:id", endpoint.EditExpenseByID)
}

func (r RouterSession) routerCreate(endpoint *expense.Endpoint) {
	r.Session.POST("/expenses", endpoint.AddExpense)
}

func serviceRouter() {
	var e RouterSession
	e.Session = echo.New()
	e.Session.Use(middleware.Logger())
	e.Session.Use(middleware.Recover())
	postgresDBClient := database.NewPostgres()
	e.Client = postgresDBClient.Client
	expensePostgresRepo := expense.NewPostgres(postgresDBClient.Client)
	expenseServiceAPI := expense.NewService(expensePostgresRepo)
	expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
	e.routerRead(expenseEndpoint)
	e.routerCreate(expenseEndpoint)
	e.routerUpdate(expenseEndpoint)
	e.gracefulShutdown()

}

func (r RouterSession) gracefulShutdown() {
	go func() {
		if err := r.Session.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			log.ErrorLog(err.Error(), "ECHO API")
			r.Session.Logger.Fatal("shutting down the server")
		}
		log.InfoLog("ECHO API START", "ECHO API")
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := r.Session.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server %s", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}

	if err := r.Client.Close(); err != nil {
		fmt.Printf("Error closing db connection %s", err)
	} else {
		fmt.Println("DB connection gracefully closed")
	}
}

func EchoStart() {
	serviceRouter()

	log.InfoLog("ECHO API STOP", "ECHO API")
}
