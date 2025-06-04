package controllers

import (
	"backend-iftar-gdgoc/models"
	"fmt"
)

func CRUDTeman(teman *[]models.Teman) {
	fmt.Print("Masukkan nama teman: ")
	var nama string
	fmt.Scanln(&nama)

	fmt.Print("Masukkan divisi teman: ")
	var divisi string
	fmt.Scanln(&divisi)

	*teman = append(*teman, models.Teman{Nama: nama, Divisi: divisi})
	fmt.Println("Teman ditambahkan:", nama, "-", divisi)
}
