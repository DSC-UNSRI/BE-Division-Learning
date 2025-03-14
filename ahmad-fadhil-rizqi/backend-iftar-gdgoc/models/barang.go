package models

import (
	"backend-iftar-gdgoc/data"
	"fmt"
)

type Barang struct {
	Nama string
}

func CRUDBarang(barang *[]Barang) {
	for {
		fmt.Print("Masukkan nama barang (atau ketik 'selesai' untuk keluar): ")
		var input string
		fmt.Scanln(&input)

		if input == "selesai" {
			break
		}

		*barang = append(*barang, Barang{Nama: input})
		data.CatatLog("Barang ditambahkan: " + input)
		fmt.Println("Barang ditambahkan:", input)
	}
}
