package account

import (
	"database/sql"
	"net/http"
	"transactionroutine/db"
	"transactionroutine/model"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	a := new(model.Account)
	a.AccountID = c.Param("accountId")

	conn := db.GetConn()
	err := FindAccount(conn, a)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Account not found")
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"account_id":      a.AccountID,
			"document_number": a.DocumentNumber})
}

func FindAccount(conn *sql.DB, a *model.Account) error {
	err := conn.QueryRow(`
		SELECT document_number
		FROM Account
		WHERE account_id = ?`,
		a.AccountID).Scan(&a.DocumentNumber)
	if err != nil {
		return err
	}
	return nil
}
