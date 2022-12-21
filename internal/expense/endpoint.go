package expense

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	ServiceExpense ServicesExpense
}

type ServicesExpense interface {
	AddExpense(stdNme string) error
}

type Request struct {
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type response struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewEndpoint(ServiceExpense ServicesExpense) *Endpoint {
	return &Endpoint{
		ServiceExpense: ServiceExpense,
	}
}

func (e Endpoint) AddExpense(c echo.Context) error {
	var r Request
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:       "Failed",
			ErrorMessage: err.Error(),
		})
	}

	fmt.Println(r.Title, r.Amount, r.Note, r.Tags)
	if err := e.ServiceExpense.AddExpense("rrrr"); err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:       "Failed",
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, response{
		Status:       "Success",
		ErrorMessage: "",
	})
}
