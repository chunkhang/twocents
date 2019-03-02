package controller

import (
	"fmt"
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// CreateSession creates a new session
func CreateSession(c echo.Context) (err error) {
	defer util.Catch(&err)

	email := c.FormValue("email")
	password := c.FormValue("password")

	fmt.Printf("email = %+v\n", email)
	fmt.Printf("password = %+v\n", password)

	return c.Redirect(http.StatusSeeOther, "/")
}
