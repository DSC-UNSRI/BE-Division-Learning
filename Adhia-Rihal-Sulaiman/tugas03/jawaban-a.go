package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	// Membaca bilangan dan menghitung frekuensi
	frekuensi := make(map[int]int)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		frekuensi[num]++
	}
	
	// Mencari modus terbesar
	var modusTerbesar, frekuensiTertinggi int
	
	for bilangan, jumlah := range frekuensi {
		if jumlah > frekuensiTertinggi || (jumlah == frekuensiTertinggi && bilangan > modusTerbesar) {
			modusTerbesar = bilangan
			frekuensiTertinggi = jumlah
		}
	}
	
	// Menampilkan hasil
	fmt.Println(modusTerbesar)
}
