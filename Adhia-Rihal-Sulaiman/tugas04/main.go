package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tugas04/controllers"
	"tugas04/models"
)

func createExplanationFile() {
	content := `Sistem Dashboard Iftar GDCoC

Sistem ini dibuat untuk mendata peserta yang mengikuti iftar GDCoC.

Fitur-fitur:
1. Autentikasi sederhana dengan email dan password
2. Manajemen opsi kendaraan (maksimal 2 dari 3 pilihan)
3. Pendataan barang yang dibawa
4. Penambahan rekomendasi untuk iftar
5. Pendataan teman yang ikut iftar
6. Melihat semua data

Sistem menggunakan godotenv untuk mengelola kredensial login.
`

	file, err := os.Create("explanation.txt")
	if err != nil {
		fmt.Println("Error membuat file explanation.txt:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error menulis ke file explanation.txt:", err)
	}
}

func main() {
	err := controllers.LoadEnv()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		fmt.Println("Pastikan file .env berada di direktori yang sama dengan aplikasi.")
		return
	}

	createExplanationFile()

	success := false
	for attempts := 0; attempts < 3 && !success; attempts++ {
		success = controllers.Login()
		if !success && attempts < 2 {
			fmt.Println("Silakan coba lagi.")
		}
	}

	if !success {
		fmt.Println("Terlalu banyak percobaan gagal. Aplikasi berhenti.")
		return
	}

	dashboard := models.Dashboard{
		Transportation:  []models.Transportation{},
		Items:           []models.Item{},
		Recommendations: []models.Recommendation{},
		Friends:         []models.Friend{},
	}

	for {
		choice := controllers.DisplayMenu()

		switch choice {
		case 1:
			controllers.ManageTransportation(&dashboard)
		case 2:
			controllers.ManageItems(&dashboard)
		case 3:
			controllers.ManageRecommendations(&dashboard)
		case 4:
			controllers.ManageFriends(&dashboard)
		case 5:
			controllers.DisplayAllData(&dashboard)
		case 6:
			fmt.Println("Terima kasih telah menggunakan Dashboard Iftar GDCoC!")
			return
		}
	}
}