package main

import (
	"fmt"
)

func carian(n int, angka []int) (int, int) {
	setar := make([]bool, n+1)

	for _, num := range angka {
		setar[num] = true
	}

	hilang := []int{}
	for i := 1; i <= n; i++ {
		if !setar[i] {
			hilang = append(hilang, i)
		}
	}

	return hilang[0], hilang[1]
}

func main() {
	var n int
	fmt.Scan(&n)

	angka := make([]int, n-2)
	for i := 0; i < n-2; i++ {
		fmt.Scan(&angka[i])
	}

	hilang1, hilang2 := carian(n, angka)
	fmt.Println(hilang1)
	fmt.Println(hilang2)
}
