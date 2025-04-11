package main

import (
	"fmt"
	"strconv"
)

func reverseInt(n int) int {
	s := strconv.Itoa(n)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedStr := string(runes)
	reversedInt, _ := strconv.Atoi(reversedStr) //Abaikan error karena input pasti valid
	return reversedInt
}
func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		fmt.Println(reverseInt(num))
	}
}