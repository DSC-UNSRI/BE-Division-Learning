package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

}
