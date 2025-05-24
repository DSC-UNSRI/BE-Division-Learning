package config

import (
	"fmt"
	"log"
	"os"
	"tugas-5/models"

	"github.com/joho/godotenv"
)

func ENVLoad() (models.User, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := models.User{
		Name:     os.Getenv("NAMA"),
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}

	if user.Email == "" || user.Password == "" {
		return models.User{}, fmt.Errorf("email or password missing in .env")
	}

	return user, nil

}
