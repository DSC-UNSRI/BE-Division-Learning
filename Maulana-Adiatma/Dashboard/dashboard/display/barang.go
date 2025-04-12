package display

import (
	"fmt"
	"cobafix.go/dashboard/controllers"
)

func TampilkanBarang() {
	var pilihan int
	for {
		fmt.Println("\nMenu Barang:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan (1-3): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			controllers.InputBarang()
		case 2:
			controllers.ViewBarang()
			controllers.HapusBarang()
		case 3:
			fmt.Println("Keluar")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi")
		}
	}
}