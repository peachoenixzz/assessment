package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	ServiceExpense ServicesExpense
}

type ServicesExpense interface {
	InsertExpense(stdNme string) error
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

func (e Endpoint) InsertExpense(c echo.Context) error {
	var request struct {
		CustomerName string `json:"customer_name"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:       "Failed",
			ErrorMessage: err.Error(),
		})
	}

	if err := e.ServiceExpense.InsertExpense(request.CustomerName); err != nil {
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
