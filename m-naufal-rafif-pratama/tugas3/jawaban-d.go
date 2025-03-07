package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	ketemu := make([]bool, N+1)

	for i := 0; i < N-2; i++ {
		var num int
		fmt.Scan(&num)
		ketemu[num] = true
	}

	hilang := []int{}
	for i := 1; i <= N; i++ {
		if !ketemu[i] {
			hilang = append(hilang, i)
		}
	}

	fmt.Println(hilang[0])
	fmt.Println(hilang[1])
}