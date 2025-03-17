package autentikasi

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Login() bool {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return false
	}

	var email, password string
	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	if  email == os.Getenv("EMAIL") && password == os.Getenv("PASSWORD") {
		fmt.Println("Login berhasil!")
		return true
	} else {
		fmt.Println("Email atau password salah")
		return false
	}
}
