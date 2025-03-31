package display

import (
	"fmt"
	"cobafix.go/dashboard/controllers"
)

func TampilkanKendaraan() {
	for {
		selected := controllers.GetKendaraan()

		
		if selected == "" {
			fmt.Println("\nkamu belum memilih kendaraan! Silakan pilih kendaraan:")
		} else {
			fmt.Println("\nkamu sudah memilih:", selected)
			fmt.Print("kamu ingin memperbarui kendaraan? (y/n): ")

			var updateChoice string
			fmt.Scanln(&updateChoice)

			if updateChoice != "y" && updateChoice != "Y" {
				fmt.Println("Pilihan kendaraan:", selected)
				return
			}
		}

		fmt.Println("\nPilih kendaraan untuk pergi ke iftar:")
		fmt.Println("1. Kendaraan Pribadi")
		fmt.Println("2. Bus Kaleng")
		fmt.Println("3. Nebeng")
		fmt.Println("4. Travel")
		fmt.Print("Masukkan pilihan (1-4): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		controllers.Kendaraan(pilihan)
		selected = controllers.GetKendaraan()

		if selected != "" {
			fmt.Println("kamu memilih:", selected)
			return 
		}

		fmt.Println("Pilihan tidak valid, coba lagi!")
	}
}
