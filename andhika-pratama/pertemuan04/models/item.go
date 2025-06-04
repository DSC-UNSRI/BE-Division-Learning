package models

import "fmt"

var items []string

func AddItem() {
    var item string
    fmt.Print("Masukkan barang yang akan dibawa: ")
    fmt.Scan(&item)
    items = append(items, item)
    fmt.Println("Barang berhasil ditambahkan.")
}
