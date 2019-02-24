package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Home is the controller for index page
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"msg": "Hello World!",
	})
}
