package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func ENVLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("Environment variables loaded successfully.")
}