package api

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/peachoenixz/assessment/internal/database"
	"github.com/peachoenixz/assessment/internal/expense"
	"github.com/peachoenixz/assessment/pkg/log"
)

<<<<<<< Updated upstream
=======
type Router interface {
	routerRead(endpoint *expense.Endpoint)
	routerUpdate(endpoint *expense.Endpoint)
	routerCreate(endpoint *expense.Endpoint)
}

type RouterSession struct {
	Session *echo.Echo
}

func (r RouterSession) routerRead(endpoint *expense.Endpoint) {
	r.Session.GET("/expenses/:id", endpoint.ViewExpenseByID)
	r.Session.GET("/expenses", endpoint.ViewExpense)
}

func (r RouterSession) routerUpdate(endpoint *expense.Endpoint) {
	r.Session.PUT("/expenses/:id", endpoint.EditExpenseByID)
}

func (r RouterSession) routerCreate(endpoint *expense.Endpoint) {
	r.Session.POST("/expenses", endpoint.AddExpense)
}

>>>>>>> Stashed changes
func serviceRouter() {
	e := echo.New()
	postgresDBClient := database.NewPostgres()
	expensePostgresRepo := expense.NewPostgres(postgresDBClient.Client)
	expenseServiceAPI := expense.NewService(expensePostgresRepo)
	expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
	e.POST("/expenses", expenseEndpoint.AddExpense)
	e.GET("/expenses/:id", expenseEndpoint.ViewExpenseByID)

	log.InfoLog("ECHO PREPARE TO START", "ECHO API")
	log.ErrorLog(e.Start(os.Getenv("PORT")), "ECHO API")
}

func EchoStart() {
	serviceRouter()
	log.InfoLog("ECHO API STOP", "ECHO API")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
