package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	counts := make(map[int]int)
	maxCount, modusTerbesar := 0, 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		counts[num]++

		switch {
		case counts[num] > maxCount:
			maxCount, modusTerbesar = counts[num], num
		case counts[num] == maxCount && num > modusTerbesar:
			modusTerbesar = num
		}
	}

	fmt.Println(modusTerbesar)
}
