//go:build unit
// +build unit

package api

import (
	"github.com/lib/pq"
	"github.com/peachoenixz/assessment/internal/expense"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestViewExpense(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow("1", "apple smoothie", 89, "no discount", pq.Array(&[]string{"beverage"}))

	db, mock, err := sqlmock.New()
	mock.ExpectPrepare("SELECT id,title,amount,note,tags FROM expenses ORDER BY id ASC").ExpectQuery().WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expensePostgresRepo := expense.NewPostgresMock(db)
	expenseServiceAPI := expense.NewService(expensePostgresRepo)
	expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
	c := e.NewContext(req, rec)
	expected := `[{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}]`

	// Act
	err = expenseEndpoint.ViewExpense(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}

func TestViewExpenseByID(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow("1", "apple smoothie", 89, "no discount", pq.Array(&[]string{"beverage"}))

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectPrepare("SELECT id,title,amount,note,tags FROM expenses WHERE id=$1").ExpectQuery().
		WithArgs("1").WillReturnRows(newsMockRows)

	expensePostgresRepo := expense.NewPostgresMock(db)
	expenseServiceAPI := expense.NewService(expensePostgresRepo)
	expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	expected := `{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}`
	// Act
	err = expenseEndpoint.ViewExpenseByID(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}

func TestAddExpenses(t *testing.T) {
	//Arrange
	ex := struct {
		ID     int      `json:"id"`
		Title  string   `json:"title"`
		Amount float64  `json:"amount"`
		Note   string   `json:"note"`
		Tags   []string `json:"tags"`
	}{
		ID:     1,
		Title:  "apple bla bla",
		Amount: 500,
		Note:   "buy apple but no discount",
		Tags:   []string{"market"},
	}

	jsonParam := `{"title":"apple bla bla","amount":500,"note":"buy apple but no discount","tags":["market"]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(string(jsonParam)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)

	newsMockRows := sqlmock.NewRows([]string{"title", "amount", "note", "tags", "id"}).
		AddRow("apple bla bla", 500, "buy apple but no discount", pq.Array(&[]string{"market"}), "1")

	mock.ExpectQuery("INSERT INTO expenses (title,amount,note,tags) values ($1,$2,$3,$4) RETURNING title,amount,note,tags,id").
		WithArgs(ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags)).WillReturnRows(newsMockRows)

	expensePostgresRepo := expense.NewPostgresMock(db)
	expenseServiceAPI := expense.NewService(expensePostgresRepo)
	expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
	c := e.NewContext(req, rec)
	expected := `{"id":1,"title":"apple bla bla","amount":500,"note":"buy apple but no discount","tags":["market"]}`
	// Act
	err = expenseEndpoint.AddExpense(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}
