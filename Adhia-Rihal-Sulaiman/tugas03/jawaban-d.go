package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	ada := make([]bool, n+1)
	
	// Membaca N-2 bilangan yang tersedia dan menandainya
	var num int
	for i := 0; i < n-2; i++ {
		fmt.Scan(&num)
		ada[num] = true
	}
	
	// Mencari dan mencetak bilangan yang hilang
	for i := 1; i <= n; i++ {
		if !ada[i] {
			fmt.Println(i)
		}
	}
}
