package display

import (
	"fmt"

	"cobafix.go/dashboard/controllers"
	"cobafix.go/dashboard/models"
)

var daftarRekomendasi []models.Rekomendasi

func TampilkanRekomendasi() {
	for {
		fmt.Println("\nMenu Rekomendasi:")
		fmt.Println("1. Tambah Rekomendasi")
		fmt.Println("2. Hapus Rekomendasi")
		fmt.Println("3. Perbarui Rekomendasi")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan (1-4): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			controllers.TambahRekomendasi(&daftarRekomendasi)
		case 2:
			controllers.LihatRekomendasi(daftarRekomendasi)
			controllers.HapusRekomendasi(&daftarRekomendasi)
		case 3:
			controllers.LihatRekomendasi(daftarRekomendasi)
			controllers.PerbaruiRekomendasi(&daftarRekomendasi)
		case 4:
			fmt.Println("Keluar")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
