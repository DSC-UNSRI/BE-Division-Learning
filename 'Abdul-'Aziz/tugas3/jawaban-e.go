package main

import "fmt"

func sumMultiples(n, divisor int) int {
	count := n / divisor
	return divisor * count * (count + 1) / 2
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)

		sum4 := sumMultiples(N, 4)
		sum7 := sumMultiples(N, 7)
		sum28 := sumMultiples(N, 28)

		fmt.Println(sum4 + sum7 - sum28)
	}
}