package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword - Menghasilkan hash untuk password
func HashPassword(password string) (string, error) {
	// Menggunakan bcrypt untuk membuat hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(bytes), nil
}

// ComparePasswordHash - Membandingkan password yang diberikan dengan hash yang ada
func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
