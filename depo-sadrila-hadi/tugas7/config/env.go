package config

import (
	"log"

	"github.com/joho/godotenv"
)

func ENVLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables if set.")
	}
}