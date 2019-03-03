package controller

import (
	"net/http"

	"github.com/labstack/echo"
	store "github.com/labstack/echo-contrib/session"
)

// HomePage returns the home page
func HomePage(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "home.tmpl", map[string]interface{}{
		"title":    "Hello World",
		"subtitle": "Welcome to Two Cents",
	})
}

// RegisterPage returns the register page
func RegisterPage(c echo.Context) (err error) {
	session, _ := store.Get("user_session", c)
	flashes := session.Flashes()
	session.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "register.tmpl", map[string]interface{}{
		"csrf":    c.Get("csrf").(string),
		"flashes": flashes,
	})
}

// LoginPage returns the login page
func LoginPage(c echo.Context) (err error) {
	session, _ := store.Get("user_session", c)
	flashes := session.Flashes()
	session.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "login.tmpl", map[string]interface{}{
		"csrf":    c.Get("csrf").(string),
		"flashes": flashes,
	})
}

// AboutPage returns the about page
func AboutPage(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "about.tmpl", nil)
}
