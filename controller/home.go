package controller

import (
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// HomePage returns the home page
func HomePage(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "home.tmpl", map[string]interface{}{
		"title":    "Hello World",
		"subtitle": "Welcome to Two Cents",
	})
	util.Check(err)

	return
}
