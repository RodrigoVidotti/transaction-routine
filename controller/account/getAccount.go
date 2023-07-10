package account

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	id := c.Param("accountId")

	resp := fmt.Sprintf("returning account %s", id)

	return c.JSON(http.StatusOK, resp)
}
