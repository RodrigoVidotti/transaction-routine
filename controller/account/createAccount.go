package account

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Account struct {
	DocumentNumber string `json:"document_number"`
}

func CreateAccount(c echo.Context) error {
	a := new(Account)
	if err := c.Bind(a); err != nil {
		return err
	}

	resp := fmt.Sprintf("account created | document_number: %s", a.DocumentNumber)

	return c.JSON(http.StatusCreated, resp)
}
