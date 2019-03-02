package server

import (
	"fmt"

	"github.com/chunkhang/twocents/config"
	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
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

	s.setConfig()

	s.setRoutes()

	err = s.setRenderer()
	util.Check(err)

	port := fmt.Sprintf(":%s", config.Port)
	s.Start(port)

	return
}
