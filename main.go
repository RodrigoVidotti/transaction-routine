package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Account struct {
	DocumentNumber string `json:"document_number"`
}

func createAccount(c echo.Context) error {
	a := new(Account)
	if err := c.Bind(a); err != nil {
		return err
	}

	resp := fmt.Sprintf("account created | document_number: %s", a.DocumentNumber)

	return c.JSON(http.StatusCreated, resp)
}

func getAccount(c echo.Context) error {
	id := c.Param("accountId")

	resp := fmt.Sprintf("returning account %s", id)

	return c.JSON(http.StatusOK, resp)
}

type transaction struct {
	AccountId       string `json:"account_id"`
	OperationTypeId string `json:"operation_type_id"`
	Amount          string `json:"amount"`
}

func createTransaction(c echo.Context) error {
	t := new(transaction)
	if err := c.Bind(t); err != nil {
		return err
	}

	resp := fmt.Sprintf("transaction created | account_id: %s | operation_type_id: %s | amount: %s", t.AccountId, t.OperationTypeId, t.Amount)
	return c.JSON(http.StatusCreated, resp)
}

func main() {
	e := echo.New()

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "He's alive!!!!\nhttps://www.youtube.com/watch?v=-5BX_q5XOHY&ab_channel=CrlosAcevegam")
	})
	e.POST("/accounts", createAccount)
	e.GET("/accounts/:accountId", getAccount)
	e.POST("/transactions", createTransaction)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
