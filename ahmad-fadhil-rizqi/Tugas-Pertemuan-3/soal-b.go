package main

import "fmt"

func main() {
	// Deklarasi variabel
	var n int
	var angka int

    // Slice untuk menyimpan hasil balik angka
	var hasil []int

	// Membaca jumlah bilangan
	fmt.Scan(&n)



	// Loop untuk membaca dan membalikkan setiap angka
	for i := 0; i < n; i++ {
		// Membaca angka
		fmt.Scan(&angka)

		// Menyimpan hasil balik angka dalam slice
		hasil = append(hasil, balikangka(angka))
	}

	// Menampilkan hasil setelah semua semua input diprses
	for _, res := range hasil {
		fmt.Println(res)
	}
}

// Fungsi untuk membalikkan angka
func balikangka(x int) int {
	var reversed int
	for x > 0 {
		// Mengambil digit terkhir
		reversed = reversed*10 + x%10
		// Hapus digit terakhir
		x = x / 10
	}
	return reversed
}
