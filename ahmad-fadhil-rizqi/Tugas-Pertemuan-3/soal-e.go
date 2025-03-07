package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var jumlahT int    
	var bilanganN int           

	// Membaca jumlah
	fmt.Scan(&jumlahT)

	// Loop untuk setiap T
	for i := 0; i < jumlahT; i++ {
		// Membaca nilai N
		fmt.Scan(&bilanganN)

		// Menghitung jumlah kelipatan 4
		k4 := bilanganN / 4
		jumlahKelipatan4 := 4 * k4 * (k4 + 1) / 2

		// Menghitung jumlah kelipatan 7
		k7 := bilanganN / 7
		jumlahKelipatan7 := 7 * k7 * (k7 + 1) / 2

		// Menghitung jumlah kelipatan 28 (kelipatan 4 dan 7)
		k28 := bilanganN / 28
		jumlahKelipatan28 := 28 * k28 * (k28 + 1) / 2

		// Menghitung total kelipatan 4 atau 7
		total := jumlahKelipatan4 + jumlahKelipatan7 - jumlahKelipatan28

		// Mencetak hasil
		fmt.Println(total)
	}
}