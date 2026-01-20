package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Username string
}

func NewEnv() (*Env, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	env := Env{
		Username: os.Getenv("username"),
	}

	return &env, nil
}
