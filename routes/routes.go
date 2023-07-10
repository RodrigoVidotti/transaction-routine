package routes

import (
	"net/http"
	"transactionroutine/controller"
	"transactionroutine/controller/account"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "He's alive!!!!\nhttps://www.youtube.com/watch?v=-5BX_q5XOHY&ab_channel=CrlosAcevegam")
	})

	accountGroup := e.Group("/account")
	{
		accountGroup.POST("", account.CreateAccount)
		accountGroup.GET("/:accountId", account.GetAccount)
	}

	e.POST("/transaction", controller.CreateTransaction)

	return e
}
