package dashboard

import "fmt"

func Menambahrekomendasi() {
	var kategori, deskripsi string

	fmt.Println("===== Tambah Rekomendasi untuk Iftar =====")
	fmt.Print("Masukkan kategori (contoh: Film, Makanan, Buku): ")
	fmt.Scan(&kategori)

	fmt.Print("Masukkan rekomendasi dalam kategori ", kategori, ": ")
	fmt.Scan(&deskripsi)

	fmt.Println("\nRekomendasi berhasil ditambahkan!")
	fmt.Println("Kategori:", kategori)
	fmt.Println("Deskripsi:", deskripsi)
}
