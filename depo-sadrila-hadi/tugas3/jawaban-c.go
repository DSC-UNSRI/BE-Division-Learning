package main

import "fmt"

func reverse(x int) int {
	temp := x
	ret := 0

	for temp > 0 {
		ret = (ret * 10) + (temp % 10)
		temp = temp / 10
	}

	return ret
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	reversedA := reverse(a)
	reversedB := reverse(b)

	sum := reversedA + reversedB
	reversedSum := reverse(sum)

	fmt.Println(reversedSum)
}