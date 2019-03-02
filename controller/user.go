package controller

import (
	"errors"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

type registerForm struct {
	Email           string `form:"email" valid:"required~Email is blank,email~Invalid email format"`
	Password        string `form:"password" valid:"required~Password is blank,length(6|15)~Password length should be 6 - 15"`
	PasswordConfirm string `form:"password_confirm" valid:"required~Confirm Password is blank"`
}

// CreateUser creates a new user
func CreateUser(c echo.Context) (err error) {
	defer util.Catch(&err)

	form := new(registerForm)
	if err = c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	_, err = valid.ValidateStruct(form)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if form.Password != form.PasswordConfirm {
		return c.String(http.StatusBadRequest, errors.New("Passwords do not match").Error())
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
