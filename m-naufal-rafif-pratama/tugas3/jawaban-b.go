package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) int {
	str := strconv.Itoa(n)
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	reversedNum, _ := strconv.Atoi(reversedStr)
	return reversedNum
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
