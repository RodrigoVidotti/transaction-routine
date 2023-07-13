package account

import (
	"database/sql"
	"net/http"
	"transactionroutine/db"
	"transactionroutine/model"

	"github.com/labstack/echo/v4"
)

func CreateAccount(c echo.Context) error {
	a := new(model.Account)
	if err := c.Bind(a); err != nil {
		panic(err.Error())
	}

	conn := db.GetConn()

	accountID, err := InsertAccount(conn, a)
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusCreated, accountID)
}

func InsertAccount(conn *sql.DB, a *model.Account) (int64, error) {
	resp, err := conn.Exec(`
		INSERT INTO Account
		(document_number)
		VALUES
		(?)`,
		a.DocumentNumber)

	if err != nil {
		panic(err.Error())
	}

	accountID, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return accountID, err
}
