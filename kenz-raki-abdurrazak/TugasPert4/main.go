package main

import (
	"fmt"
	"github.com/artichys/BE-Division-Learning/config"
	"github.com/artichys/BE-Division-Learning/services"
	"github.com/artichys/BE-Division-Learning/utils"
)

func main() {
	config.LoadEnv()

	if !services.Authenticate() {
		return
	}

	for {
		fmt.Println("\nDashboard Iftar GDGoC")
		fmt.Println("1. Pilih Kendaraan (1 opsi)")
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Tambah Rekomendasi")
		fmt.Println("4. Tambah Teman")
		fmt.Println("5. Lihat Data")
		fmt.Println("6. Exit")
		fmt.Print("Pilih menu: ")

		choice := utils.ReadInt() // 🔥 Ganti fmt.Scan dengan ReadInt()

		switch choice {
		case 1:
			services.ChooseVehicle()
		case 2:
			services.AddItem()
		case 3:
			services.AddRecommendation()
		case 4:
			services.AddFriend()
		case 5:
			services.ShowData()
		case 6:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
