package main

import "fmt"

// Fungsi untuk membalikkan angka
func reverse(x int) int {
	temp := x
	ret := 0

	for temp > 0 {
		ret = (ret * 10) + (temp % 10)
		temp = temp / 10
	}

	return ret
}

func main() {
	// Deklarasi variabel
	var a int
	var b int

	// Membaca input
	fmt.Scan(&a, &b)

	// Membalikkan angka a dan b, menjumlahkan, kemudian membalikkan hasil penjumlahannya
	result := reverse(reverse(a) + reverse(b))

	// Menampilkan hasil
	fmt.Println(result)
}
