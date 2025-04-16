package main

import (
	"fmt"
)

// Func for reversing the integers
func reverseInt(num int) int {
	var rev int

	// Check if the current num higher than 0 (this will process all the 'individual integer' of each index within num until it reaches 0)
	// If its not and its in the middle of the 'integer string', it will simply return the 0
	// If its a consecutive last zeros in the og integer, it'll simply multiply with itself and add itself with another 0, resulting in not a single 0 being in the beginning of the rev integer (unless the input is just consesutive 0s) even if there are multiple 0s in the end of the og integer
	for num > 0 { 					
		rev = rev * 10 + num % 10
		num /= 10
	}

	return rev
}	

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		fmt.Println(reverseInt(num)) // Call the reverseInt function to iterate each num input-ed and print it immediately
	}
}