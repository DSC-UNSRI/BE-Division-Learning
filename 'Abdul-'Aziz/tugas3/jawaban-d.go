package main

import "fmt"

func main() {
	var n int
	fmt.Println("Masukkan panjang sequence:")
	fmt.Scan(&n)

	// Menggunakan map untuk tracking angka yang sudah ada
	numbers := make(map[int]bool)
	fmt.Printf("Masukkan %d angka:\n", n-2)
	for i := 0; i < n-2; i++ {
		var num int
		fmt.Scan(&num)
		numbers[num] = true
	}

	// Karena kita tahu hanya ada 2 angka yang hilang
	// dan angka berurutan dari 1 sampai n
	// kita bisa langsung print saat menemukannya
	fmt.Println("Bilangan yang hilang:")
	found := 0
	for i := 1; i <= n && found < 2; i++ {
		if !numbers[i] {
			fmt.Println(i)
			found++
		}
	}
}