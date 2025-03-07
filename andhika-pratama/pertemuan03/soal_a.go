package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n) 					

	numFreq := make(map[int]int)	// Map that store the freqs of each nums

	maxFreq := 0 
	maxModus := 0

	for i := 0; i < n; i++ { 
		var num int
		fmt.Scan(&num)
		numFreq[num]++				// Update the freqs of each nums

		// If the current freq num is higher than maxFreq, update the maxFreq and maxModus
		// If the current freq num is equal to the maxFreq and the current num is higher maxModus, update the maxFreq(tho meaningless) and maxModus
		if maxFreq < numFreq[num]  || (maxFreq == numFreq[num] && maxModus < num) {
			maxFreq = numFreq[num]
			maxModus = num
		}
	}
	
	fmt.Println(maxModus)
}	