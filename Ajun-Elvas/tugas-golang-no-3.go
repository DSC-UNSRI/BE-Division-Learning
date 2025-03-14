package main

import (
	"fmt"
)

func tukar(x int) int {
	k := x
	m := 0

	for k > 0 {
		m = (m * 10) + (k % 10)
		k = k / 10
	}

	return m
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	atukar := tukar(a)
	btukar := tukar(b)

	c := atukar + btukar
	ctukar := tukar(c)

	fmt.Println(ctukar)
}
