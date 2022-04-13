package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/product"
	"be/delivery/controllers/user"
	"be/delivery/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo, ac *auth.Controller, uc *user.Controller, pc *product.Controller) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	// e.AcquireContext().Cookies()
	/* no jwt */

	e.GET("/test", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})

	// login ====================================

	e.POST("/login", ac.Login())

	// user ====================================

	e.POST("/user", uc.Create())

	/* with jwt */

	var g = e.Group("")

	g.Use(middlewares.JwtMiddleware())

	// product ====================================

	g.POST("/product", pc.Create())
	g.GET("/product", pc.Get())
	g.PUT("/product", pc.Update())
	g.DELETE("/product", pc.Delete())
}
