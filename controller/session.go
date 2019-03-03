package controller

import (
	"net/http"
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	store "github.com/labstack/echo-contrib/session"
)

type loginForm struct {
	Email    string `form:"email" valid:"required~Email is blank,email~Invalid email format"`
	Password string `form:"password" valid:"required~Password is blank,length(6|15)~Password length should be 6 - 15"`
}

// CreateSession creates a new session for user
func CreateSession(c echo.Context) (err error) {
	session, _ := store.Get("user_session", c)

	form := new(loginForm)
	if err = c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	_, err = valid.ValidateStruct(form)
	if err != nil {
		messages := strings.Split(err.Error(), ";")
		for _, message := range messages {
			session.AddFlash(message)
		}
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/login")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
