package controller

import (
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// Register returns the register page
func Register(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "register.tmpl", map[string]interface{}{
		"title":    "Hello World",
		"subtitle": "Welcome to Two Cents",
	})
	util.Check(err)

	return
}
