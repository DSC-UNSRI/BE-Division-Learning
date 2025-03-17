package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	freq := make([]int, 1001)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		freq[num]++
	}

	maxFreq := 0
	maxNum := 0

	for num, count := range freq {
		if count > maxFreq || (count == maxFreq && num > maxNum) {
			maxFreq = count
			maxNum = num
		}
	}

	fmt.Println(maxNum)
}