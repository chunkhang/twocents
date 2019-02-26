package controller

import (
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// Home is the controller for index page
func Home(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
		"msg": "Hello World!",
	})
	util.Check(err)

	return
}
