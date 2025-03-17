package dashboard

import "fmt"

func Menambahbarang() {
	fmt.Println("======= Tambah Barang ======")
	fmt.Println("= !Dilarang keras membawa! =")
	fmt.Println("=    Senjata Api & Tajam   =")
	fmt.Println("============================")
	
	var namaBarang string
	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&namaBarang)
	fmt.Println(namaBarang, "berhasil ditambahkan!")
}
