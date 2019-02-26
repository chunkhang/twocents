package route

import (
	"github.com/chunkhang/twocents/controller"
	"github.com/chunkhang/twocents/util"
	"github.com/chunkhang/twocents/view"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init initalizes the router
func Init() (e *echo.Echo, err error) {
	defer util.Catch(&err)
	e = echo.New()

	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	renderer, err := view.Init()
	util.Check(err)

	e.Renderer = renderer

	e.GET("/", controller.Home)

	return
}
