package article

import (
	"echo-framework/article/controllers"
	"echo-framework/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	g := e.Group("/article")
	g.Use(middleware.Auth)
	g.GET("", controllers.GetListArticle)
	g.GET("/:id", controllers.GetArticleByID)
	g.POST("", controllers.CreateArticle)
	g.PUT("", controllers.UpdateArticle)
	g.DELETE("", controllers.DeleteArticle)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
}
