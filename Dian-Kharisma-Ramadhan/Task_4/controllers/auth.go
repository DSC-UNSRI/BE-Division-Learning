package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	"Task_4/models"
)

func LoadEnv() models.User {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return models.User{
		Name:     os.Getenv("NAMA"),
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}
}
func Autentikasi(user models.User) bool {
	var email, password string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Print("Masukkan Password: ")
	scanner.Scan()
	password = scanner.Text()

	return email == user.Email && password == user.Password
}
