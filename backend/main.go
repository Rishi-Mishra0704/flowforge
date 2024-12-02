package main

import (
	"log"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
	"github.com/Rishi-Mishra0704/flowforge/backend/server"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	server, err := server.NewServer(config)

	if err != nil {
		log.Fatal(err)
	}
	server.Start(config.PORT)

}
