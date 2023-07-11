package account

import (
	"net/http"
	"transactionroutine/db"
	"transactionroutine/model"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	acc := new(model.Account)
	acc.AccountID = c.Param("accountId")

	conn := db.GetConn()
	err := conn.QueryRow(`
		SELECT document_number
		FROM Account
		WHERE account_id = ?`,
		acc.AccountID).Scan(&acc.DocumentNumber)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Account not found")
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"account_id":      acc.AccountID,
			"document_number": acc.DocumentNumber})
}
