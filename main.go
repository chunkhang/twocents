package main

import (
	"github.com/chunkhang/twocents/server"
	"github.com/chunkhang/twocents/util"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	err := server.Start()
	util.Check(err)
}
