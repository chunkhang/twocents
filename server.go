package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Configuration
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n",
	}))
	e.Use(middleware.Recover())
	portNumber := "3001"

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", portNumber)))
}
