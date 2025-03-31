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
		fmt.Println("3. Pilih Rekomendasi")
		fmt.Println("4. Lihat semua data")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan pilihan (1-5): ")

		var menu int
		fmt.Scan(&menu)

		switch menu {
		case 1:
			display.TampilkanKendaraan()
		case 2:
			display.TampilkanBarang()
		case 3:
			display.TampilkanRekomendasi()
		case 4:
			display.LihatSemuaData()
		case 5:
			fmt.Println("Terima kasih! Keluar dari program.")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid, coba lagi!")
		}
	}
}
