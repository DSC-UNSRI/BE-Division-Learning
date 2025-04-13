package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Login() bool {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Gagal load file .env")
		return false
	}

	envEmail := os.Getenv("EMAIL")
	envPassword := os.Getenv("PASSWORD")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Email: ")
	inputEmail, _ := reader.ReadString('\n')
	inputEmail = strings.TrimSpace(inputEmail)

	fmt.Print("Password: ")
	inputPassword, _ := reader.ReadString('\n')
	inputPassword = strings.TrimSpace(inputPassword)

	return inputEmail == envEmail && inputPassword == envPassword
}
