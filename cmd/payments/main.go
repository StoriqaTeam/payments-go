package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/storiqateam/payments/app"
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	app.PrintConfig()
	app.StartServer()
}
