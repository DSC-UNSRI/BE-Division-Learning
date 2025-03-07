package main

import (
	"fmt"
)

func sumKelipatan(n, x int) int {
	k := n / x
	return x * (k * (k + 1) / 2)
}

func main() {
	var jumlahPertanyaan int
	fmt.Scan(&jumlahPertanyaan)

	for i := 0; i < jumlahPertanyaan; i++ {
		var batasBilangan int
		fmt.Scan(&batasBilangan)

		total := sumKelipatan(batasBilangan, 4) + sumKelipatan(batasBilangan, 7) - sumKelipatan(batasBilangan, 28)

		fmt.Println(total)
	}
}
