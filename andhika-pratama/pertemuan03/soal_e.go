package main

import "fmt"


func sumMultiples(n, num int) int {
	m := n / num						// Calculates the number of multiples that num has up to n
	sum := m * (m + 1) / 2				// Calculate all the sum of the first natural numbers from 1 up to m (arithmetic formula for counting the sum of the first natural numbers)
	return num * sum					// Return said sum * num coz if u basically just multiple the natural numbers above with the current num we are trying to find the sum of the multiples for, well get the sum for the later
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		// Calculate the sum of multiples of 4, 7, and 28 up to n
		// Note that we count the sum for 28 coz its the least common multiple of 4 and 7
		sum4 := sumMultiples(n, 4)
		sum7 := sumMultiples(n, 7)
		sum28 := sumMultiples(n, 28)

		// We minus the totalSum to sum28 coz its the numbers that will be double counted (both counted in 4 and 7)
		totalSum := sum4 + sum7 - sum28

		fmt.Println(totalSum)
	}
}
