package controller

import (
	"net/http"
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	store "github.com/labstack/echo-contrib/session"
)

type registerForm struct {
	Email           string `form:"email" valid:"required~Email is blank,email~Invalid email format"`
	Password        string `form:"password" valid:"required~Password is blank,length(6|15)~Password length should be 6 - 15"`
	PasswordConfirm string `form:"password_confirm" valid:"required~Confirm Password is blank"`
}

// CreateUser creates a new user
func CreateUser(c echo.Context) (err error) {
	session, _ := store.Get("user_session", c)

	form := new(registerForm)
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

		return c.Redirect(http.StatusSeeOther, "/register")
	}

	if form.Password != form.PasswordConfirm {
		session.AddFlash("Passwords do not match")
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/register")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
