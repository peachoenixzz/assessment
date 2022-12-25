//go:build integration
// +build integration

package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/peachoenixz/assessment/internal/expense"
	"github.com/stretchr/testify/assert"
)

const serverPort = 2565

type response struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func getLastID(res response, db *sql.DB) (response, error) {
	stmt, err := db.Prepare("select distinct on (id) id from expenses order by id desc;")
	if err != nil {
		log.Fatal(err)
		return response{}, err
	}
	err = stmt.QueryRow().Scan(&res.ID)
	if err != nil {
		return response{}, err
	}
	return res, nil
}

func SetAddExpense(jsonString string, db *sql.DB) (string, error) {
	var res response
	if err := json.Unmarshal([]byte(jsonString), &res); err != nil {
		return "", err
	}
	res, err := getLastID(res, db)
	if err != nil {
		return "", err
	}
	ret, err := json.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func postgresqlService() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://root:root@db/go-db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func startService() (*echo.Echo, *sql.DB) {
	eh := echo.New()
	db := postgresqlService()
	go func(e *echo.Echo) {
		expensePostgresRepo := expense.NewPostgres(db)
		expenseServiceAPI := expense.NewService(expensePostgresRepo)
		expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
		e.GET("/expenses/:id", expenseEndpoint.ViewExpenseByID)
		e.GET("/expenses", expenseEndpoint.ViewExpense)
		e.POST("/expenses", expenseEndpoint.AddExpense)
		e.PUT("/expenses/:id", expenseEndpoint.EditExpenseByID)
		e.Start(fmt.Sprintf(":%d", serverPort))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
	return eh, db
}

func TestViewExpense(t *testing.T) {
	// Setup server
	eh, _ := startService()
	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `[{"id":1,"title":"buy a new phone","amount":39000,"note":"buy a new phone","tags":["gadget","shopping"]}]`

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestViewExpenseByID(t *testing.T) {
	// Setup server
	eh, _ := startService()
	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/1", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `{"id":1,"title":"buy a new phone","amount":39000,"note":"buy a new phone","tags":["gadget","shopping"]}`

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestAddExpense(t *testing.T) {
	// Setup server
	eh, _ := startService()
	// Arrange

	jsonString := `{"title":"Edit : apple not smoothie","amount":150,"note":"have discount","tags":["market"]}`

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(jsonString))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `{"id":2,"title":"Edit : apple not smoothie","amount":150,"note":"have discount","tags":["market"]}`
	if err != nil {
		t.Fatal(err)
	}
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestUpdateExpenseByID(t *testing.T) {
	// Setup server
	eh, _ := startService()
	// Arrange

	jsonString := `{"title":"Really we got 1 update","amount":500,"note":"update is success","tags":["market","success update"]}`

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d/expenses/1", serverPort), strings.NewReader(jsonString))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `{"id":1,"title":"Really we got 1 update","amount":500,"note":"update is success","tags":["market","success update"]}`

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
