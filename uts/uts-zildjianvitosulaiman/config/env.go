package config

import (
	"log"
	"uts-zildjianvitosulaiman/models"

	"github.com/joho/godotenv"
)

func ENVLoad() (models.User, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
