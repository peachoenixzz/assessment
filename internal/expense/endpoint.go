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
	AddExpense(req Request) (int, error)
}

type Request struct {
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Response struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Amount       float64  `json:"amount"`
	Note         string   `json:"note"`
	Tags         []string `json:"tags"`
	Status       string   `json:"status"`
	ErrorMessage string   `json:"error_message"`
}

func NewEndpoint(ServiceExpense ServicesExpense) *Endpoint {
	return &Endpoint{
		ServiceExpense: ServiceExpense,
	}
}

func (e Endpoint) AddExpense(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:       "Failed",
			ErrorMessage: err.Error(),
		})
	}

	fmt.Println(req.Title, req.Amount, req.Note, req.Tags)
	id, err := e.ServiceExpense.AddExpense(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:       "Failed",
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, Response{
		ID:           id,
		Title:        req.Title,
		Note:         req.Note,
		Amount:       req.Amount,
		Tags:         req.Tags,
		Status:       fmt.Sprintf("%v", http.StatusCreated),
		ErrorMessage: "",
	})
}
