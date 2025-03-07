package main

import (
	"fmt"
)

func findMissingNumbers(N int, numbers []int) (int, int) {
	present := make([]bool, N+1)
	for _, num := range numbers {
		present[num] = true
	}

	missing := []int{}
	for i := 1; i <= N; i++ {
		if !present[i] {
			missing = append(missing, i)
		}
	}
	return missing[0], missing[1]
}

func MainFindMissingNumbers() {
	var N int
	fmt.Scan(&N)

	numbers := make([]int, N-2)
	for i := 0; i < N-2; i++ {
		fmt.Scan(&numbers[i])
	}

	missing1, missing2 := findMissingNumbers(N, numbers)
	fmt.Println(missing1)
	fmt.Println(missing2)
}
