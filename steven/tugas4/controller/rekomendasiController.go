package controller

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
	"tugas4/models"
)

func RekomendasiController(dashboard *models.Dashboard){
	fmt.Println("Pilih Kategori")
	fmt.Println("1. Film")
	fmt.Println("2. Game")
	fmt.Println("3. Kembali Ke Dashboard")

	var opsi int;
	fmt.Print("Masukkan Opsi :")
	fmt.Scan(&opsi)
	fmt.Scanln()

	var kategori string
	var rekomendasi string
	switch opsi {
	case 1 :
		kategori = "Film"
		fmt.Print("Masukkan Rekomendasi Film :")
	case 2 : 
		kategori = "Game"
		fmt.Print("Masukkan Rekomendasi Game :")
	case 3 :
		Dashboard(dashboard)
	default :
		fmt.Println("Opsi tidak valid, coba lagi.")
		RekomendasiController(dashboard)
	}

	reader := bufio.NewReader(os.Stdin)
	rekomendasi, _ = reader.ReadString('\n')
	rekomendasi = strings.TrimSpace(rekomendasi)

	if dashboard.Rekomendasi == nil {
		dashboard.Rekomendasi = make(map[string][]string)
	}
	
	dashboard.Rekomendasi[kategori] = append(dashboard.Rekomendasi[kategori], rekomendasi)

	fmt.Println("Rekomendasi berhasil ditambahkan!")
	RekomendasiController(dashboard)
}