package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make(map[int]bool)
	for i := 0; i < n-2; i++ {
		var num int
		fmt.Scan(&num)
		numbers[num] = true
	}

	missing := []int{}
	for i := 1; i <= n; i++ {
		if !numbers[i] {
			missing = append(missing, i)
		}
	}

	sort.Ints(missing) // Urutkan untuk memastikan yang lebih kecil dicetak duluan

	fmt.Println(missing[0])
	fmt.Println(missing[1])
}