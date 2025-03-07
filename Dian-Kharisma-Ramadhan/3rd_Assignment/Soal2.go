package main

import "fmt"

func terbalik(n int) int {
	reverse := 0
	for n > 0 {
		lastdigit := n % 10
		reverse = reverse*10 + lastdigit
		n /= 10
	}
	return reverse
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var angka int
		fmt.Scan(&angka)
		fmt.Println(terbalik(angka))
	}
}
