package controllers

import (
	"fmt"
)

type Rekomendasi struct {
	Kategori string
	Isi      string
}

func CRUDRekomendasi(rekomendasi *[]Rekomendasi) {
	fmt.Print("Masukkan kategori (misalnya: Film, Buku, Musik): ")
	var kategori string
	fmt.Scanln(&kategori)

	fmt.Print("Masukkan rekomendasi: ")
	var isi string
	fmt.Scanln(&isi)

	*rekomendasi = append(*rekomendasi, Rekomendasi{Kategori: kategori, Isi: isi})
	fmt.Println("Rekomendasi ditambahkan:", kategori, "-", isi)
}
