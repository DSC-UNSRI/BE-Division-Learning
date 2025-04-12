package controller

import (
	"fmt"

	"tugas4/models"
)

func PilihKendaraan(dashboard *models.Dashboard){

	fmt.Println("Pilih Kendaraan ke Iftar:")
	fmt.Println("1. Kendaraan Pribadi")
	fmt.Println("2. Bus Kaleng")
	fmt.Println("3. Nebeng")
	fmt.Println("4. Travel")

	var opsi int
	fmt.Print("Masukkan Opsi Anda :")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		dashboard.Kendaraan = "Kendaraan Pribadi"
		Dashboard(dashboard)
		fmt.Println("Kendaraan Berhasil Dipilih")
	case 2:
		dashboard.Kendaraan = "Bus Kaleng"
		Dashboard(dashboard)
		fmt.Println("Kendaraan Berhasil Dipilih")
	case 3:
		dashboard.Kendaraan = "Nebeng"
		Dashboard(dashboard)
		fmt.Println("Kendaraan Berhasil Dipilih")
	case 4:
		dashboard.Kendaraan = "Travel"
		Dashboard(dashboard)
		fmt.Println("Kendaraan Berhasil Dipilih")
	default:
		fmt.Println("Opsi tidak valid, coba lagi.")
		PilihKendaraan(dashboard)
	}
}