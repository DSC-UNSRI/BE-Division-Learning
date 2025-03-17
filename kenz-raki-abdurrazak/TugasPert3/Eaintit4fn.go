package main

import (
	"fmt"
)

func sumMultiples(n, k int) int {
	m := n / k
	return k * m * (m + 1) / 2
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		sum := sumMultiples(n, 4) + sumMultiples(n, 7) - sumMultiples(n, 28)
		fmt.Println(sum)
	}
}
