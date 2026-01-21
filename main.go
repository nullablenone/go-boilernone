package main

import (
	"log"

	"github.com/nullablenone/go-boilernone/config"
	"github.com/nullablenone/go-boilernone/routes"
)

func main() {
	_, err := config.NewEnv()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	router := routes.SetRoutes()
	if err := router.Run(":3333"); err != nil {
		log.Fatalf("failed to start HTTP server on :3333: %v", err)
	}

}
