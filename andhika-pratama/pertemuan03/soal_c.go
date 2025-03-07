package main

import "fmt"

func reverseInt(num int) int {
	var rev int

	// I dont have to comment it no more dont i?
	for num > 0 { 					
		rev = rev * 10 + num % 10
		num /= 10
	}

	return rev
}	

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	// Reverse A and B.
	revA := reverseInt(A)
	revB := reverseInt(B)

	sum := revA + revB 			// Add the reversed integers together

	revSum := reverseInt(sum) // Reverse sum

	// Print the final result.
	fmt.Println(revSum)
}
