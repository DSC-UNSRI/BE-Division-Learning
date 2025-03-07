package main

import "fmt"

func MainBiggestModus() {
	var n int
	fmt.Scan(&n)

	counts := make(map[int]int)
	maxCount, biggestModus := 0, 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		counts[num]++

		switch {
		case counts[num] > maxCount:
			maxCount, biggestModus = counts[num], num
		case counts[num] == maxCount && num > biggestModus:
			biggestModus = num
		}
	}

	fmt.Println(biggestModus)
}
