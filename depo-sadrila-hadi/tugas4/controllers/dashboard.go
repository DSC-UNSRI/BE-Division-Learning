package controllers

import (
	"bufio"
	"fmt"
	"os"
	"tugas4/models"
	"strings"
)

func ShowMenu(data *models.Data) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("\nDashboard Iftar GDGoC")
        fmt.Println("1. Pilih Kendaraan")
        fmt.Println("2. Input Barang Dibawa")
        fmt.Println("3. Input Rekomendasi")
        fmt.Println("4. Data Teman")
        fmt.Println("5. Lihat Semua Data")
        fmt.Println("6. Keluar")
        fmt.Print("Pilih opsi: ")
        scanner.Scan()
        option := scanner.Text()

        switch option {
        case "1":
            fmt.Println("Pilih Kendaraan (Kendaraan Pribadi, Bus Kaleng, Nebeng, Travel): ")
            scanner.Scan()
            kendaraan := scanner.Text()
            data.Kendaraan = []string{kendaraan}
        case "2":
            fmt.Println("Input barang (pisahkan koma jika banyak): ")
            scanner.Scan()
            barang := strings.Split(scanner.Text(), ",")
            data.Barang = barang
        case "3":
            fmt.Println("Kategori rekomendasi: ")
            scanner.Scan()
            kategori := scanner.Text()
            fmt.Println("Isi rekomendasi (pisahkan koma jika banyak): ")
            scanner.Scan()
            isi := strings.Split(scanner.Text(), ",")
            data.Rekomendasi[kategori] = isi
        case "4":
            fmt.Println("Nama teman: ")
            scanner.Scan()
            nama := scanner.Text()
            fmt.Println("Divisi teman: ")
            scanner.Scan()
            divisi := scanner.Text()
            data.Teman = append(data.Teman, models.Teman{Nama: nama, Divisi: divisi})
        case "5":
            fmt.Printf("Kendaraan: %v\n", data.Kendaraan)
            fmt.Printf("Barang: %v\n", data.Barang)
            fmt.Printf("Rekomendasi: %v\n", data.Rekomendasi)
            fmt.Printf("Teman: %v\n", data.Teman)
        case "6":
            fmt.Println("Keluar Dashboard")
            return
        default:
            fmt.Println("Opsi tidak valid")
        }
    }
}