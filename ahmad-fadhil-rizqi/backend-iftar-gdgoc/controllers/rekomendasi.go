package controllers

import (
	"backend-iftar-gdgoc/models"
	"fmt"
)



func CRUDRekomendasi(rekomendasi *[]models.Rekomendasi) {
	fmt.Print("Masukkan kategori (misalnya: Film, Buku, Musik): ")
	var kategori string
	fmt.Scanln(&kategori)

	fmt.Print("Masukkan rekomendasi: ")
	var isi string
	fmt.Scanln(&isi)

	*rekomendasi = append(*rekomendasi, models.Rekomendasi{Kategori: kategori, Isi: isi})
	fmt.Println("Rekomendasi ditambahkan:", kategori, "-", isi)
}
