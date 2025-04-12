package display

import (
	"fmt"

	"cobafix.go/dashboard/controllers"
	"cobafix.go/dashboard/models"
)

var daftarTeman []models.Teman

func TampilkanTeman() {
	for {
		fmt.Println("\nMenu Teman:")
		fmt.Println("1. Tambah Teman")
		fmt.Println("2. Hapus Teman")
		fmt.Println("3. Perbarui Teman")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan (1-4): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			controllers.TambahReTeman(&daftarTeman)
		case 2:
			controllers.ViewTeman(daftarTeman)
			controllers.HapusTeman(&daftarTeman)
		case 3:
			controllers.ViewTeman(daftarTeman)
			controllers.PerbaruiTeman(&daftarTeman)
		case 4:
			fmt.Println("Keluar")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi")
		}
	}
}
