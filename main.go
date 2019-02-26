package main

import (
	"fmt"

	"github.com/chunkhang/twocents/config"
	"github.com/chunkhang/twocents/route"
	"github.com/chunkhang/twocents/util"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	router, err := route.Init()
	util.Check(err)

	port := fmt.Sprintf(":%s", config.Port)
	router.Start(port)
}
