package main

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/chunkhang/twocents/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Configuration
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.HideBanner = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n",
	}))
	e.Use(middleware.Recover())

	// Templates
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("view/*.tmpl")),
	}

	// Routes
	e.GET("/", handler.HomeHandler)

	// Start server
	port, ok := os.LookupEnv("PORT")
	if !ok {
		e.Logger.Fatal("PORT environment variable was not set")
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
