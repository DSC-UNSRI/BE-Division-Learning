package dashboard

import (
	"fmt"
	"tugasday4/dashboard/function"
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

		switch pilihan {
		case 1:
			dashboard.Pilihkendaraan()
		case 2:
			dashboard.Menambahbarang()
		case 6:
			fmt.Println("Keluar dari dashboard...")
			return
		default:
			fmt.Println("Pilihan mu tidak sesuai silahkan input 1 - 6")
		}
	}
}
