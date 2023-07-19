package account

import (
	"database/sql"
	"net/http"
	"strconv"
	"transactionroutine/db"
	"transactionroutine/model"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	var err error

	a := new(model.Account)
	a.AccountID, err = strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Account not found")
	}

	conn := db.GetConn()
	err = FindAccount(conn, a)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Account not found")
	}

	return c.JSON(
		http.StatusOK,
		map[string]int{
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
