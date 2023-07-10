package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transaction struct {
	AccountId       string `json:"account_id"`
	OperationTypeId string `json:"operation_type_id"`
	Amount          string `json:"amount"`
}

func CreateTransaction(c echo.Context) error {
	t := new(transaction)
	if err := c.Bind(t); err != nil {
		return err
	}

	resp := fmt.Sprintf("transaction created | account_id: %s | operation_type_id: %s | amount: %s", t.AccountId, t.OperationTypeId, t.Amount)
	return c.JSON(http.StatusCreated, resp)
}
