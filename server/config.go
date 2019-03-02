package server

import (
	"fmt"

	"github.com/labstack/echo/middleware"
)

const (
	staticDirectory = "static"
	assetDirectory  = "assets"
	logFormat       = "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n"
)

func (s *server) setConfig() {
	s.HideBanner = true
	s.Static(fmt.Sprintf("/%s", staticDirectory), assetDirectory)
	s.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat,
	}))
	s.Use(middleware.Gzip())
	s.Use(middleware.CORS())
	s.Use(middleware.Recover())
}
