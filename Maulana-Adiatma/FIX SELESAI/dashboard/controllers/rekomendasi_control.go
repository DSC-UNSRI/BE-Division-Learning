package controllers

import (
	"fmt"

	"cobafix.go/dashboard/models"
)

func TambahRekomendasi(daftarRekomendasi *[]models.Rekomendasi) {
	var kategori, isi string
	fmt.Print("Masukkan kategori (hanya satu kata): ")
	fmt.Scan(&kategori)
	fmt.Print("Masukkan isi rekomendasi (hanya satu kata): ")
	fmt.Scan(&isi)

	*daftarRekomendasi = append(*daftarRekomendasi, models.Rekomendasi{Kategori: kategori, Isi: isi})
	fmt.Println("Rekomendasi berhasil ditambahkan!")
}

func HapusRekomendasi(daftarRekomendasi *[]models.Rekomendasi) {
	var isi string
	fmt.Print("Masukkan isi rekomendasi yang ingin dihapus: ")
	fmt.Scan(&isi)

	for i, r := range *daftarRekomendasi {
		if r.Isi == isi {
			*daftarRekomendasi = append((*daftarRekomendasi)[:i], (*daftarRekomendasi)[i+1:]...)
			fmt.Println("Rekomendasi berhasil dihapus!")
			return
		}
	}
	fmt.Println("Rekomendasi tidak ditemukan.")
}

func PerbaruiRekomendasi(daftarRekomendasi *[]models.Rekomendasi) {
	var kategoriLama, isiLama, kategoriBaru, isiBaru string
	fmt.Print("Masukkan kategori : ")
	fmt.Scan(&kategoriLama)
	fmt.Print("Masukkan isi : ")
	fmt.Scan(&isiLama)

	for i, r := range *daftarRekomendasi {
		if r.Isi == isiLama && r.Kategori == kategoriLama {
			fmt.Println("Rekomendasi ditemukan:", r.Kategori, "-", r.Isi)
			fmt.Print("Masukkan kategori baru: ")
			fmt.Scan(&kategoriBaru)
			fmt.Print("Masukkan isi baru: ")
			fmt.Scan(&isiBaru)

			(*daftarRekomendasi)[i] = models.Rekomendasi{Kategori: kategoriBaru, Isi: isiBaru}
			fmt.Println("Rekomendasi berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Rekomendasi tidak ditemukan.")
}

func ViewRekomendasi(daftarRekomendasi []models.Rekomendasi) {
	if len(daftarRekomendasi) == 0 {
		fmt.Println("Kamu belum input rekomendasi.")
		return
	}

	for _, r := range daftarRekomendasi {
		fmt.Println("-", r.Kategori, ":", r.Isi)
	}
}
