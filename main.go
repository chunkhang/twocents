package main

import (
	"fmt"

	"github.com/chunkhang/twocents/config"
	"github.com/chunkhang/twocents/route"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	router := route.Init()

	port := fmt.Sprintf(":%s", config.Port)
	router.Start(port)
}
