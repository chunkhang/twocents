package server

import (
	"fmt"

	"github.com/chunkhang/twocents/config"
	"github.com/chunkhang/twocents/controller"
	"github.com/chunkhang/twocents/util"
	"github.com/chunkhang/twocents/view"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	staticDirectory = "static"
	assetDirectory  = "assets"
	logFormat       = "${time_rfc3339}: ${method} ${uri} ${status} (${latency_human})\n"
)

type server struct {
	*echo.Echo
}

// Start initializes and starts the server
func Start() (err error) {
	defer util.Catch(&err)

	s := &server{
		echo.New(),
	}

	s.config()
	s.routes()
	err = s.renderer()
	util.Check(err)

	s.Start(fmt.Sprintf(":%s", config.Port))

	return
}

func (s *server) config() {
	s.HideBanner = true
	s.Static(fmt.Sprintf("/%s", staticDirectory), assetDirectory)
	s.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat,
	}))
	s.Use(middleware.Gzip())
	s.Use(middleware.CORS())
	s.Use(middleware.Recover())
}

func (s *server) routes() {
	s.GET("/", controller.HomePage)
	s.GET("/register", controller.RegisterPage)
	s.GET("/login", controller.LoginPage)
	s.GET("/about", controller.AboutPage)
	s.POST("/users", controller.CreateUser)
}

func (s *server) renderer() (err error) {
	defer util.Catch(&err)

	r, err := view.Init()
	util.Check(err)

	s.Renderer = r

	return
}
