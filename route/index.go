package route

import (
	"github.com/chunkhang/twocents/handler"
	"github.com/chunkhang/twocents/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init initalizes the router
func Init() *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.Renderer = template.Init()

	e.GET("/", handler.HomeHandler)

	return e
}
