package controller

import (
	"net/http"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

// HomePage returns the home page
func HomePage(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "home.tmpl", map[string]interface{}{
		"title":    "Hello World",
		"subtitle": "Welcome to Two Cents",
	})
	util.Check(err)

	return
}

// RegisterPage returns the register page
func RegisterPage(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "register.tmpl", map[string]interface{}{
		"csrf": c.Get("csrf").(string),
	})
	util.Check(err)

	return
}

// LoginPage returns the login page
func LoginPage(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "login.tmpl", map[string]interface{}{
		"csrf": c.Get("csrf").(string),
	})
	util.Check(err)

	return
}

// AboutPage returns the about page
func AboutPage(c echo.Context) (err error) {
	defer util.Catch(&err)

	err = c.Render(http.StatusOK, "about.tmpl", nil)
	util.Check(err)

	return
}
