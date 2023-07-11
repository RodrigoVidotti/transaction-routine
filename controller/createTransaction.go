package controller

import (
	"net/http"
	"transactionroutine/db"
	"transactionroutine/model"

	"github.com/labstack/echo/v4"
)

func CreateTransaction(c echo.Context) error {
	t := new(model.Transaction)
	if err := c.Bind(t); err != nil {
		return err
	}

	conn := db.GetConn()
	resp, err := conn.Exec(`
		INSERT INTO Transaction
		(account_id, operation_type_id, amount)
		VALUES
		(?,?,?)`,
		t.AccountId, t.OperationTypeId, t.Amount)

	if err != nil {
		panic(err.Error())
	}

	transactionID, err := resp.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusCreated, transactionID)
}
