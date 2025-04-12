package main

import "fmt"

func reverse(temp int) int {
	reversed := 0
	for temp > 0 {
		reversed = (reversed * 10) + (temp % 10)
		temp /= 10
	}
	return reversed
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	A_reversed := reverse(A)
	B_reversed := reverse(B)

	C := A_reversed + B_reversed

	fmt.Println(reverse(C))
}
