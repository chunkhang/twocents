package server

import (
	"fmt"

	"github.com/chunkhang/twocents/app/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

const (
	staticDirectory = "static"
	assetDirectory  = "assets"
	logFormat       = "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n"
	cookiePath      = "/"
	cookieHTTPOnly  = true
)

func (s *server) setConfig() {
	s.HideBanner = true

	s.Static(fmt.Sprintf("/%s", staticDirectory), assetDirectory)

	s.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat,
	}))

	s.Use(middleware.Gzip())

	s.Use(middleware.CORS())

	s.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:csrf",
		CookiePath:     cookiePath,
		CookieHTTPOnly: cookieHTTPOnly,
	}))

	s.Use(middleware.Secure())

	s.Use(middleware.Recover())

	store := sessions.NewCookieStore(config.StoreKey)
	store.Options = &sessions.Options{
		HttpOnly: cookieHTTPOnly,
	}
	s.Use(session.Middleware(store))
}
