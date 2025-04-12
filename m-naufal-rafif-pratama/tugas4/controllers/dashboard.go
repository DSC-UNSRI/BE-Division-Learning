package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tugas4/models"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func StartDashboard() {
	data := models.NewData()

	for {
		fmt.Println("\nDashboard Iftar GDGoC")
		fmt.Println("1. Pilih Kendaraan")
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Tambah Rekomendasi")
		fmt.Println("4. Tambah Teman")
		fmt.Println("5. Lihat Data")
		fmt.Println("6. Exit")

		choice := getInput("Pilih menu: ")

		switch choice {
		case "1":
			data.Vehicle = getInput("Pilih kendaraan (Kendaraan Pribadi/Bus Kaleng/Nebeng): ")
		case "2":
			item := getInput("Masukkan barang yang akan dibawa: ")
			data.Items = append(data.Items, item)
		case "3":
			category := getInput("Masukkan kategori rekomendasi: ")
			rec := getInput("Masukkan isi rekomendasi: ")
			data.Recommendations[category] = rec
		case "4":
			friendName := getInput("Masukkan nama teman: ")
			division := getInput("Masukkan divisi teman: ")
			data.Friends[friendName] = division
		case "5":
			fmt.Println("\nData Anda:")
			fmt.Println("Kendaraan:", data.Vehicle)
			fmt.Println("Barang yang dibawa:", data.Items)
			fmt.Println("Rekomendasi:", data.Recommendations)
			fmt.Println("Teman yang ikut:", data.Friends)
		case "6":
			fmt.Println("Keluar dari sistem")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
