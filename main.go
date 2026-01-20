package main

import (
	"fmt"
	"log"

	"github.com/nullablenone/go-boilernone/config"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	fmt.Println(env.Username)
}
