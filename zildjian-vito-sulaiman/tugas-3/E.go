package main

import (
	"fmt"
)

func sumDivisibleBy(x, N int) int {
	m := N / x
	return x * m * (m + 1) / 2
}

func sumMultiples(N int) int {
	return sumDivisibleBy(4, N) + sumDivisibleBy(7, N) - sumDivisibleBy(28, N)
}

func MainSumMultiples() {
	var T, N int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Println(sumMultiples(N))
	}
}
