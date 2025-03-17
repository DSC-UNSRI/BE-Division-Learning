package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		
		hasil := hitungJumlahKelipatan(n)
		fmt.Println(hasil)
	}
}

func hitungJumlahKelipatan(n int) int {
	kelipatan4 := n / 4
	sum4 := 4 * kelipatan4 * (kelipatan4 + 1) / 2
	
	kelipatan7 := n / 7
	sum7 := 7 * kelipatan7 * (kelipatan7 + 1) / 2
	
	kelipatan28 := n / 28
	sum28 := 28 * kelipatan28 * (kelipatan28 + 1) / 2
	
	return sum4 + sum7 - sum28
}
