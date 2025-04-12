package utils

import (
	"fmt"
	"tugas-5/models"
)

func Authenticate(user models.User) bool {
	var email, password string

	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	return email == user.Email && password == user.Password
}
