package main

import (
	"github.com/storiqateam/payments/app"
)

func main() {
	app.PrintConfig()
	app.StartServer()
}
