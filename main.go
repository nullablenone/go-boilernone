package main

import (
	"log"

	"github.com/nullablenone/go-boilernone/config"
	"github.com/nullablenone/go-boilernone/routes"
)

func main() {
	_, err := config.ConnectPostgre()
	if err != nil {
		log.Fatal(err)
	}

	router := routes.SetRoutes()
	if err = router.Run(":3333"); err != nil {
		log.Fatalf("failed to start HTTP server on :3333: %v", err)
	}

}
