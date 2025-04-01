package controllers

import (
	"fmt"

	"cobafix.go/dashboard/models"
)

func TambahReTeman(daftarTeman *[]models.Teman) {
	var nama, divisi string
	fmt.Print("Masukkan nama teman (hanya satu kata): ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan divisi teman (hanya satu kata): ")
	fmt.Scan(&divisi)

	*daftarTeman = append(*daftarTeman, models.Teman{Nama: nama, Divisi: divisi})
	fmt.Println("Teman berhasil diajak!")
}

func HapusTeman(daftarTeman *[]models.Teman) {
	var nama, divisi string
	fmt.Print("Masukkan nama teman (hanya satu kata): ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan divisi teman (hanya satu kata): ")
	fmt.Scan(&divisi)

	for i, r := range *daftarTeman {
		if r.Nama == nama && r.Divisi == divisi {
			*daftarTeman = append((*daftarTeman)[:i], (*daftarTeman)[i+1:]...)
			fmt.Println("Teman berhasil dihapus!")
			return
		}
	}
	fmt.Println("Teman tidak ditemukan.")
}

func PerbaruiTeman(daftarTeman *[]models.Teman) {
	var namaLama, divisiLama, namaBaru, divisiBaru string
	fmt.Print("Masukkan Nama Teman: ")
	fmt.Scan(&namaLama)
	fmt.Print("Masukkan Divisi Teman: ")
	fmt.Scan(&divisiLama)

	for i, r := range *daftarTeman {
		if r.Nama == namaLama && r.Divisi == divisiLama {
			fmt.Println("Rekomendasi ditemukan:", r.Nama, "-", r.Divisi)
			fmt.Print("Masukkan Update Nama Teman: ")
			fmt.Scan(&namaBaru)
			fmt.Print("Masukkan Update Divisi Teman: ")
			fmt.Scan(&divisiBaru)

			(*daftarTeman)[i] = models.Teman{Nama: namaBaru, Divisi: divisiBaru}
			fmt.Println("Temanmu berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Teman tidak ditemukan.")
}

func ViewTeman(daftarTeman []models.Teman) {
	if len(daftarTeman) == 0 {
		fmt.Println("Kamu belum input teman.")
		return
	}

	for _, r := range daftarTeman {
		fmt.Println("-", r.Nama, ":", r.Divisi)
	}
}
