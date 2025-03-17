package dashboard

import (
	"fmt"
)

func MenuDashboard() {
	for {
		fmt.Println("=== Dashboard ===")
		fmt.Println("1. Pilih Kendaraan")
		fmt.Println("2. Tambah Barang Bawaan")
		fmt.Println("3. Tambah Rekomendasi Kegiatan")
		fmt.Println("4. Ajak Teman")
		fmt.Println("5. Lihat Semua Data")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)
	}
}
