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
	envNama := os.Getenv("USERNAMES")
	envEmail := os.Getenv("EMAIL")
	envPassword := os.Getenv("PAS SWORD")
	fmt.Println("Dari .env:", envNama, envEmail, envPassword)
	
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Pas sword: ")
	fmt.Scanln(&password)

	if nama == envNama && email == envEmail && password == envPassword {
		fmt.Println("Login berhasil!")
		return true
	}

	fmt.Println("Login gagal! Coba lagi")
	return false
}
