package dashboard

import "fmt"

type Teman struct {
	Nama   string
	Divisi string
}

var daftarTeman []Teman

func Menambahteman() {
	var nama, divisi string

	fmt.Println("=== Tambah Teman yang Ikut Iftar ===")
	fmt.Print("Masukkan nama teman: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan divisi teman: ")
	fmt.Scan(&divisi)

	temanBaru := Teman{Nama: nama, Divisi: divisi}
	daftarTeman = append(daftarTeman, temanBaru)
	fmt.Println("Teman berhasil ditambahkan!")
}
