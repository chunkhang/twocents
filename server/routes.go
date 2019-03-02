package server

import (
	"github.com/chunkhang/twocents/controller"
)

func (s *server) setRoutes() {
	s.GET("/", controller.HomePage)
	s.GET("/register", controller.RegisterPage)
	s.GET("/login", controller.LoginPage)
	s.GET("/about", controller.AboutPage)

	s.POST("/login", controller.CreateSession)

	s.POST("/users", controller.CreateUser)
}
