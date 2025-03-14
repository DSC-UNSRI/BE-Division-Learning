package models

import (
	"backend-iftar-gdgoc/data"
	"fmt"
)

type Teman struct {
	Nama   string
	Divisi string
}

func CRUDTeman(teman *[]Teman) {
	fmt.Print("Masukkan nama teman: ")
	var nama string
	fmt.Scanln(&nama)

	fmt.Print("Masukkan divisi teman: ")
	var divisi string
	fmt.Scanln(&divisi)

	*teman = append(*teman, Teman{Nama: nama, Divisi: divisi})
	data.CatatLog("Teman ditambahkan: " + nama + " - " + divisi)
	fmt.Println("Teman ditambahkan:", nama, "-", divisi)
}
