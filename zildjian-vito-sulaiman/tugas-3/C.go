package main

import (
	"fmt"
)

func reverseSum(x int) int {
	ret := 0
	for x > 0 {
		ret = (ret * 10) + (x % 10)
		x /= 10
	}
	return ret
}

func MainReverseSum() {
	var A, B int
	fmt.Scan(&A, &B)

	A_reversed := reverseSum(A)
	B_reversed := reverseSum(B)

	C := A_reversed + B_reversed
	C_reversed := reverseSum(C)

	fmt.Println(C_reversed)
}
