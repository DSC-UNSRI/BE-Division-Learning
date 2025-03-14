package main

import "fmt"

func hitungKelipatan(x, n int) int {
	m := n / x
	return x * m * (m + 1) / 2
}

func main() {
	var T, N int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		sum := hitungKelipatan(4, N) + hitungKelipatan(7, N) - hitungKelipatan(28, N)
		fmt.Println(sum)
	}
}
