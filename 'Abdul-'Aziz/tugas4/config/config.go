package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Name     string
	Email    string
	Password   string
)
func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Name = os.Getenv("NAME")
	Email = os.Getenv("EMAIL")
	Password = os.Getenv("PASSWORD")
}