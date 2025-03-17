package dashboard

import "fmt"

var barang []string
var rekomendasi = make(map[string]string)
var daftarTeman []Teman

func LihatData() {
	fmt.Println("\n=== Semua Data yang Anda Masukkan ===")

	if kendaraan1 == "" && kendaraan2 == "" {
		fmt.Println("Kendaraan: (Belum dipilih)")
	} else {
		fmt.Println("Kendaraan 1:", kendaraan1)
		fmt.Println("Kendaraan 2:", kendaraan2)
	}

	fmt.Println("\nBarang yang Dibawa:")
	if len(barang) == 0 {
		fmt.Println("(Belum ada barang)")
	} else {
		for i, item := range barang {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}

	fmt.Println("\nRekomendasi:")
	if len(rekomendasi) == 0 {
		fmt.Println("(Belum ada rekomendasi)")
	} else {
		for kategori, isi := range rekomendasi {
			fmt.Printf("- %s: %s\n", kategori, isi)
		}
	}

	fmt.Println("\nTeman yang Ikut Iftar:")
	if len(daftarTeman) == 0 {
		fmt.Println("(Belum ada teman yang ditambahkan)")
	} else {
		for i, teman := range daftarTeman {
			fmt.Printf("%d. Nama: %s | Divisi: %s\n", i+1, teman.Nama, teman.Divisi)
		}
	}
}
