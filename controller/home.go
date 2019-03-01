package controller

import (
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// Home returns the home page
func Home(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "home.tmpl", map[string]interface{}{
		"title":    "Hello World",
		"subtitle": "Welcome to Two Cents",
	})
	util.Check(err)

	return
}
