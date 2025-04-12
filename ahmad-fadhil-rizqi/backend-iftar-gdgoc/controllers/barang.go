package controllers

import (
	"backend-iftar-gdgoc/models"
	"fmt"
)

func CRUDBarang(barang *[]models.Barang) {
	for {
		fmt.Print("Masukkan nama barang (atau ketik 'selesai' untuk keluar): ")
		var input string
		fmt.Scanln(&input)

		if input == "selesai" {
			break
		}

		*barang = append(*barang, models.Barang{Nama: input})
		fmt.Println("Barang ditambahkan:", input)
	}
}
