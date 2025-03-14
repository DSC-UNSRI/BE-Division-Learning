package controller

import (
	"os"
	"log"
	"fmt"

	"github.com/joho/godotenv"
)

func Auth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	emailEnv := os.Getenv("EMAIL")
	passswordEnv := os.Getenv("PASSWORD")

	var email string;
	var password string;

	fmt.Println("Login Dashboard Iftar GDGoC")
	fmt.Print("EMAIL: ")
	fmt.Scan(&email)
	fmt.Print("PASSWORD: ")
	fmt.Scan(&password)

	if email == emailEnv && password == passswordEnv {
		fmt.Println("berhasil login")
	} else {
		fmt.Println("gagal login")
	}
}