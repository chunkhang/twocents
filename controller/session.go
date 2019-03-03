package controller

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
	store "github.com/labstack/echo-contrib/session"
)

type loginForm struct {
	Email    string `form:"email" valid:"required~Email is blank,email~Invalid email format"`
	Password string `form:"password" valid:"required~Password is blank,length(6|15)~Password length should be 6 - 15"`
}

// CreateSession creates a new session
func CreateSession(c echo.Context) (err error) {
	defer util.Catch(&err)

	session, _ := store.Get("session", c)

	form := new(loginForm)
	if err = c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	_, err = valid.ValidateStruct(form)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	session.Values["foo"] = "bar"
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/")
}
