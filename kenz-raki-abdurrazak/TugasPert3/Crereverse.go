package main

import (
	"fmt"
)

func reverse(x int) int {
	reversed := 0
	for x > 0 {
		reversed = (reversed * 10) + (x % 10)
		x /= 10
	}
	return reversed
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	A_reversed := reverse(A)
	B_reversed := reverse(B)
	C := A_reversed + B_reversed
	C_reversed := reverse(C)

	fmt.Println(C_reversed)
}
