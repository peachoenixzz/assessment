package expense

import (
	"github.com/peachoenixz/assessment/pkg/log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	Service ServiceUseCase
}

type ServiceUseCase interface {
	AddExpense(req Request) (Response, error)
	ViewExpense() ([]Response, error)
	ViewExpenseByID(id string) (Response, error)
	EditExpenseByID(req Request, id string) (Response, error)
}

type Request struct {
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Response struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Errors struct {
	Status  int
	Message string
}

func NewEndpoint(ServiceExpense ServiceUseCase) *Endpoint {
	return &Endpoint{
		Service: ServiceExpense,
	}
}

func (e Endpoint) AddExpense(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		log.ErrorLog(err, "bind Add expense")
		return c.JSON(http.StatusBadRequest, 400)
	}

	Response, err := e.Service.AddExpense(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Errors{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, Response)
}

func (e Endpoint) ViewExpenseByID(c echo.Context) error {
	id := c.Param("id")
	Response, err := e.Service.ViewExpenseByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Errors{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response)
}

func (e Endpoint) ViewExpense(c echo.Context) error {
	Response, err := e.Service.ViewExpense()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Errors{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response)
}

func (e Endpoint) EditExpenseByID(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, 400)
	}
	id := c.Param("id")

	Response, err := e.Service.EditExpenseByID(req, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Errors{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response)
}
