package utils

import "fmt"

// Fungsi menangani error agar tidak panic
func TanganiError(err error, pesan string) {
	if err != nil {
		fmt.Println("Error:", pesan)
	}
}
