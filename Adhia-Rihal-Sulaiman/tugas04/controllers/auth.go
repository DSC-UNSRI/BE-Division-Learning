package controllers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func CheckAuth(email, password string) bool {
	envEmail := os.Getenv("EMAIL")
	envPassword := os.Getenv("PASSWORD")

	return email == envEmail && password == envPassword
}

func Login() bool {
	var email, password string

	fmt.Println("=== Login Dashboard Iftar GDCoC ===")
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if CheckAuth(email, password) {
		fmt.Println("Login berhasil!")
		return true
	} else {
		fmt.Println("Email atau password salah!")
		return false
	}
}