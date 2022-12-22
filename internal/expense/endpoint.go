package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	Service ServiceUseCase
}

type ServiceUseCase interface {
	AddExpense(req Request) (int, error)
	ViewExpense(id string) (Response, error)
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
		return c.JSON(http.StatusBadRequest, 400)
	}

	id, err := e.Service.AddExpense(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, 500)
	}

	return c.JSON(http.StatusCreated, Response{
		ID:     id,
		Title:  req.Title,
		Note:   req.Note,
		Amount: req.Amount,
		Tags:   req.Tags,
	})
}

func (e Endpoint) ViewExpense(c echo.Context) error {
	id := c.Param("id")
	Response, err := e.Service.ViewExpense(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Errors{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, Response)
}
