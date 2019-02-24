package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/chunkhang/twocents/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
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

	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.tmpl", "view/base.tmpl"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
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
