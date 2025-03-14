package controllers

import (
	"backend-iftar-gdgoc/data"
	"backend-iftar-gdgoc/models"
	"fmt"
)

var (
	vehicleList     []models.Vehicle
	daftarBarang     []models.Barang
	daftarRekomendasi []models.Rekomendasi
	daftarTeman      []models.Teman
)

func Dashboard() {
	for {
		fmt.Println("\n===== Dashboard Iftar GDGoC =====")
		fmt.Println("1. Select Vehicle")
		fmt.Println("2. Input Barang")
		fmt.Println("3. Tambah Rekomendasi Iftar")
		fmt.Println("4. Tambah Teman yang Ikut")
		fmt.Println("5. Lihat Semua Data")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			models.SelectVehicle(&vehicleList)
		case 2:
			models.CRUDBarang(&daftarBarang)
		case 3:
			models.CRUDRekomendasi(&daftarRekomendasi)
		case 4:
			models.CRUDTeman(&daftarTeman)
		case 5:
			TampilkanData()
		case 6:
			fmt.Println("Terima kasih! Sampai jumpa di Iftar GDGoC!")
			data.CatatLog("User keluar dari sistem")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func TampilkanData() {
	fmt.Println("\n===== Data Iftar GDGoC =====")
	fmt.Println("\nSelected Vehicle:")
	for _, v := range vehicleList {
		fmt.Println("- " + v.Name)
	}
	fmt.Println("\nBarang yang Dibawa:")
	for _, b := range daftarBarang {
		fmt.Println("- " + b.Nama)
	}
	fmt.Println("\nRekomendasi Iftar:")
	for _, r := range daftarRekomendasi {
		fmt.Printf("- %s: %s\n", r.Kategori, r.Isi)
	}
	fmt.Println("\nTeman yang Ikut:")
	for _, t := range daftarTeman {
		fmt.Printf("- %s (%s)\n", t.Nama, t.Divisi)
	}
	data.CatatLog("User melihat semua data")
}
