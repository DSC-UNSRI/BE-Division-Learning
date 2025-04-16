package controllers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"tugas-4/models"
)

func LoadEnv() (models.User, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return models.User{}, fmt.Errorf("failed to load .env file")
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

func Authenticate(user models.User) bool {
	var email, password string

	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	return email == user.Email && password == user.Password
}
