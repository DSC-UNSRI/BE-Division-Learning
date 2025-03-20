package controllers

import (
	"backend-iftar-gdgoc/config"
	"fmt"
)

// Fungsi untuk meminta input dari user
func AmbilVariabel(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

// Fungsi autentikasi berdasarkan input user
func AuthenticateUser() bool {
	// Ambil email dan password dari .env
	storedEmail := config.AmbilVariabel("EMAIL")
	storedPassword := config.AmbilVariabel("PASSWORD")

	// User memasukkan email dan password
	inputEmail := AmbilVariabel("Masukkan email: ")
	inputPassword := AmbilVariabel("Masukkan password: ")

	// Validasi email dan password
	if inputEmail != storedEmail || inputPassword != storedPassword {
		fmt.Println("❌ Autentikasi gagal! Email atau password salah.")
		return false
	}

	fmt.Println("✅ Autentikasi berhasil! Selamat datang,", config.AmbilVariabel("NAMA"))
	return true
}
