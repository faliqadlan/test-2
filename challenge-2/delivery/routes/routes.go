package routes

import (
	"be/delivery/controllers/movie"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo, mc *movie.Controller) {
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

	// movie ====================================
	e.POST("/movie", mc.Create())
	e.PUT("/movie", mc.Update())
	e.DELETE("/movie", mc.Delete())
	e.GET("/movie", mc.Get())
}
