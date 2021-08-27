package routes

import (
	"linkaja/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	g := e.Group("/account")
	g.GET("/:account_number", controllers.GetAccountByNumber)
	g.POST("/:from_account_number/transfer", controllers.TransferBalance)

}
