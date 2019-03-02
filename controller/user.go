package controller

import (
	"fmt"
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// CreateUser creates a new user
func CreateUser(c echo.Context) (err error) {
	defer util.Catch(&err)

	email := c.FormValue("email")
	password := c.FormValue("password")
	passwordConfirm := c.FormValue("password_confirm")

	fmt.Printf("email = %+v\n", email)
	fmt.Printf("password = %+v\n", password)
	fmt.Printf("password_confirm = %+v\n", passwordConfirm)

	return c.Redirect(http.StatusSeeOther, "/")
}
