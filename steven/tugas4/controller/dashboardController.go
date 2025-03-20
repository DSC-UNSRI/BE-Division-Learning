package controller

import (
	"fmt"
	"log"
	"os"

	"tugas4/models"

	"github.com/joho/godotenv"
)

func Dashboard(dashboard *models.Dashboard) {	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	namaEnv := os.Getenv("NAMA")
	fmt.Println("******Welcome to Iftar GDGoC Dashboard******,", namaEnv)
	fmt.Println("Menu:")
	fmt.Println("1. Pilih Kendaraan")
	fmt.Println("2. Tambah Barang Bawaan")
	fmt.Println("3. Tambah Rekomendasi Iftar")
	fmt.Println("4. Tambah Teman")
	fmt.Println("5. Lihat Data")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih opsi: ")

	var opsi int;
	fmt.Scan(&opsi)

	switch opsi {
	case 1: 
		PilihKendaraan(dashboard)

	case 5:
		fmt.Println("===== Data Dashboard =====")
		fmt.Println("Kendaraan:", dashboard.Kendaraan)
		fmt.Println("Barang:", dashboard.Barang)
		fmt.Println("Rekomendasi:", dashboard.Rekomendasi)
		fmt.Println("Teman:", dashboard.Teman)
		Dashboard(dashboard)
	case 6:
		fmt.Println("Keluar dari program...")
		os.Exit(0)
	default:
		fmt.Println("Opsi tidak tersedia")
		Dashboard(dashboard)
	}
}