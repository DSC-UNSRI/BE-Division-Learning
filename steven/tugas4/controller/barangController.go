package controller

import (
	"fmt"
	"tugas4/models"
)

func BarangController(dashboard *models.Dashboard){
	var barang string
    fmt.Print("Masukkan barang bawaan: ")
    fmt.Scan(&barang)

    dashboard.Barang = append(dashboard.Barang, barang)
    fmt.Println("Barang berhasil ditambahkan!")
    Dashboard(dashboard)
}