package display

import (
	"fmt"

	"cobafix.go/dashboard/controllers"
)

func TampilkanKendaraan() {
	fmt.Println("Pilih kendaraan untuk pergi ke iftar:")
	fmt.Println("1. Kendaraan Pribadi")
	fmt.Println("2. Bus Kaleng")
	fmt.Println("3. Nebeng")
	fmt.Println("4. Travel")
	fmt.Print("Masukkan pilihan (1-4): ")

	var pilihan int
	fmt.Scanln(&pilihan)

	controllers.Kendaraan(pilihan) 
	selected := controllers.GetKendaraan()

	fmt.Println("Anda memilih:", selected)
}
