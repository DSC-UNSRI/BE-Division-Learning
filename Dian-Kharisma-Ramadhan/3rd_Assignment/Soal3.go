package main

import "fmt"

func terbalik(n int) int {
	reverse := 0
	for n > 0 {
		lastdigit := n % 10
		reverse = reverse*10 + lastdigit
		n /= 10
	}
	return reverse
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	// Balik angka A dan B
	A_reversed := terbalik(A)
	B_reversed := terbalik(B)

	// Jumlahkan hasil balik
	C := A_reversed + B_reversed

	// Balik hasil penjumlahan dan cetak
	fmt.Println(terbalik(C))
}
