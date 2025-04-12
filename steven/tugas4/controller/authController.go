package controller

import (
	"fmt"
	"log"
	"os"
	"tugas4/models"

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
		fmt.Println("Berhasil Login")

		dashboard := models.Dashboard{}
		Dashboard(&dashboard) 
	} else {
		fmt.Println("Gagal Login, EMAIL atau PASSWORD SALAH")
		Auth()
	}
}