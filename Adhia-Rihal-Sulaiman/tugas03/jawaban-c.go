package main

import "fmt"

func reverse(bilangan int) int {
	var hasil int
	
	for bilangan > 0 {
		digitTerakhir := bilangan % 10
		hasil = hasil*10 + digitTerakhir
		bilangan /= 10
	}
	
	return hasil
}

func main() {
	var bilanganA, bilanganB int
	fmt.Scan(&bilanganA, &bilanganB)
	
	bilanganAbalik := reverse(bilanganA)
	bilanganBbalik := reverse(bilanganB)
	

	hasilPenjumlahan := bilanganAbalik + bilanganBbalik
	
	// Langkah 3: Membalik hasil penjumlahan
	hasilAkhir := reverse(hasilPenjumlahan)
	
	// Menampilkan hasil
	fmt.Println(hasilAkhir)
}
