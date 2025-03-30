package autentikasi

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Login() bool {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Gagal membaca file .env:", err)
		return false
	}

	var nama, email, password string
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	envNama := os.Getenv("NAMA")
	envEmail := os.Getenv("EMAIL")
	envPassword := os.Getenv("PASSWORD")
  
	if nama == envNama && email == envEmail && password == envPassword {
		fmt.Println("Login berhasil!")
		return true
	}

	fmt.Println("Login gagal! Coba lagi")
	return false
}
