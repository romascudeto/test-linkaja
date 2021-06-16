package author

import (
	"echo-framework/author/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	g := e.Group("/author")
	g.GET("", controllers.GetListAuthor)
	g.GET("/:id", controllers.GetAuthorByID)
	g.POST("", controllers.CreateAuthor)
	g.PUT("", controllers.UpdateAuthor)
	g.DELETE("", controllers.DeleteAuthor)
	g.POST("/login", controllers.LoginAuthor)
}
