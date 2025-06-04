package main

import (
	"fmt"
	"os"
	"project-root/controllers"
)

func main() {
	if !controllers.Login() {
		fmt.Println("Login gagal. Keluar dari program.")
		os.Exit(1)
	}

	nama := os.Getenv("NAMA")
	email := os.Getenv("EMAIL")
	controllers.InitUserData(nama, email)

	for {
		fmt.Println("\n=== Dashboard Iftar ===")
		fmt.Println("1. Pilih Kendaraan")
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Tambah Rekomendasi")
		fmt.Println("4. Tambah Teman")
		fmt.Println("5. Lihat Semua Data")
		fmt.Println("6. Keluar")

		var pilihan int
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			controllers.PilihKendaraan()
		case 2:
			controllers.TambahBarang()
		case 3:
			controllers.TambahRekomendasi()
		case 4:
			controllers.TambahTeman()
		case 5:
			controllers.LihatData()
		case 6:
			fmt.Println("Keluar dari dashboard.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
