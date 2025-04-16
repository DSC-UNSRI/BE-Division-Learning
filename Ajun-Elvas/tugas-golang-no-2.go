package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) int {
	ang := strconv.Itoa(n)
	balik := ""

	for i := len(ang) - 1; i >= 0; i-- {
		balik += string(ang[i])
	}

	bolak, _ := strconv.Atoi(balik)
	return bolak
}

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		fmt.Println(reverseNumber(num))
	}
}
