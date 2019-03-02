package controller

import (
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// RegisterPage returns the register page
func RegisterPage(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "register.tmpl", nil)
	util.Check(err)

	return
}
