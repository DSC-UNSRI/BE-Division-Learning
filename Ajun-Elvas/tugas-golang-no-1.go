package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	counts := make(map[int]int)

	var num int
	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		counts[num]++
	}

	maxFreq := 0
	maxNum := 0

	for key, freq := range counts {
		if freq > maxFreq || (freq == maxFreq && key > maxNum) {
			maxFreq = freq
			maxNum = key
		}
	}

	fmt.Println(maxNum)
}
