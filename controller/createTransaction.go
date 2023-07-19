package controller

import (
	"database/sql"
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

	errMsg := validateAmount(t.Amount, t.OperationTypeId)
	if errMsg != nil {
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	transactionID, err := InsertTransaction(conn, t)
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusCreated, transactionID)
}

func validateAmount(amount float64, operationTypeID int) error {
	if amount == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Amount cannot be zero")
	}

	validSignal, ok := model.ValidOperationTypes[operationTypeID]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid operation type")
	}

	if amount*float64(validSignal) < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Amount signal is not compatible with the operation type")
	}

	return nil
}

func InsertTransaction(conn *sql.DB, t *model.Transaction) (int64, error) {
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
		return 0, err
	}

	return transactionID, err
}
