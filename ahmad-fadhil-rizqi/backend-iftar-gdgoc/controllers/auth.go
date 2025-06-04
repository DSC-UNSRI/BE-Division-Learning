package controllers

import (
	"fmt"
	"os"
)

func AmbilVariabel(key string) string {
	return os.Getenv(key)
}

func AuthenticateUser() bool {
	storedEmail := AmbilVariabel("EMAIL")       
	storedPassword := AmbilVariabel("PASSWORD") 

	var inputEmail, inputPassword string

	fmt.Print("Masukkan email: ")
	fmt.Scanln(&inputEmail)

	fmt.Print("Masukkan password: ")
	fmt.Scanln(&inputPassword)

	if inputEmail != storedEmail || inputPassword != storedPassword {
		fmt.Println("Autentikasi gagal! Email atau password salah.")
		return false
	}

	fmt.Println("Autentikasi berhasil! Selamat datang,", AmbilVariabel("NAMA"))
	return true


}
