package main

import (
	"fmt"
	"os"

	"github.com/chunkhang/twocents/route"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	router := route.Init()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable was not set")
	}

	router.Start(fmt.Sprintf(":%s", port))
}
