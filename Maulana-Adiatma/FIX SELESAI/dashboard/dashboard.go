package dashboard

import (
	"fmt"
	
	"cobafix.go/dashboard/display"
)

func Dashboard() {
	for {
		fmt.Println("\n===== Pilihan Menu =====")
		fmt.Println("1. Pilih kendaraan")
		fmt.Println("2. Pilih Barang")
		fmt.Println("3. Lihat semua data")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan (1-4): ")

		var menu int
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			display.TampilkanKendaraan()
		case 2:

		case 3:
			display.LihatSemuaData()
		case 4:
			fmt.Println("Terima kasih! Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}
