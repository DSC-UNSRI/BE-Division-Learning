package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	// Memproses sebanyak n bilangan
	for i := 0; i < n; i++ {
		var bilangan int
		fmt.Scan(&bilangan)
		
		// Membalik bilangan
		var hasil int
		for bilangan > 0 {
			digit := bilangan % 10
			hasil = hasil*10 + digit
			bilangan /= 10
		}
		
		fmt.Println(hasil)
	}
}
