package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	count := make([]int, 1001)

	var num int
	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		count[num]++
	}

	maxCount := 0
	modus := 0
	for i := 0; i <= 1000; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			modus = i
		} else if count[i] == maxCount && i > modus {
			modus = i
		}
	}

	fmt.Println(modus)
}
