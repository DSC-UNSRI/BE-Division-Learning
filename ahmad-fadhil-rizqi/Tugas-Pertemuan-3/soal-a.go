package main

import "fmt"

func main() {
	// Variabel yang digunakan
	var n int
	var angka int
	var maxfrekuensi int
	var modus int

	count := make(map[int]int)

	// Membaca angka yang dimasukkan
	fmt.Scan(&n)

	// Membaca angka yang diinput dan menghitung frekuensi kemunculannya
	for i := 0; i < n; i++ {
		fmt.Scan(&angka)
		count[angka]++

		// Mengupdate modus kalau frekuensinya lebih tinggi
		// Lalu jika ada angka yang lebih besar maka akan diambil modus dengan angka yang lebih besar
		if count[angka] > maxfrekuensi || (count[angka] == maxfrekuensi && angka > modus) {
			maxfrekuensi = count[angka]
			modus = angka
		}
	}

	// Menampilkan modus
	fmt.Println(modus)
}
