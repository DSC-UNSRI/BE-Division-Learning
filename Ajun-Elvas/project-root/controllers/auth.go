package controllers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
}

func Authenticate(email, password string) bool {
	LoadEnv()
	envEmail := os.Getenv("EMAIL")
	envPassword := os.Getenv("PASSWORD")

	return email == envEmail && password == envPassword
}
