package controllers

import (
    "fmt"
    "pertemuan04/models"
)

func Dashboard() {
    var choice int

    for {
        fmt.Println("\nDashboard Menu:")
        fmt.Println("1. Pilih kendaraan")
        fmt.Println("2. Tambah barang bawaan")
        fmt.Println("3. Tambah rekomendasi")
        fmt.Println("4. Tambah teman")
        fmt.Println("5. Lihat semua data")
        fmt.Println("6. Exit")
        fmt.Print("Pilih opsi: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            models.ChooseVehicle()
        case 2:
            models.AddItem()
        case 3:
            models.AddRecommendation()
        case 4:
            models.AddFriend()
        case 5:
            models.ViewAll()
        case 6:
            fmt.Println("Keluar...")
            return
        default:
            fmt.Println("Pilihan tidak valid, coba lagi.")
        }
    }
}
