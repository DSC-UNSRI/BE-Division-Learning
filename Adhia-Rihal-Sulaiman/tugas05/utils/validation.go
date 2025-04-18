package utils

import (
	"regexp"
)

// ValidasiUsername - Memastikan username hanya mengandung huruf dan angka dan panjangnya minimal 3 karakter.
func ValidasiUsername(username string) bool {
	if len(username) < 3 {
		return false
	}

	// Menggunakan regex untuk memastikan hanya huruf dan angka
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return re.MatchString(username)
}

// ValidasiEmail - Memastikan email valid
func ValidasiEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidasiPassword - Memastikan password lebih dari 6 karakter
func ValidasiPassword(password string) bool {
	return len(password) > 6
}
