package app

import (
	"fmt"
	"log"
)

// PrintConfig prints current config
func PrintConfig() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %+v", err)
	}
	fmt.Printf("%+v\n", config)
}
