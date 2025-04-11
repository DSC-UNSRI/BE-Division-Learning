package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	exist := make([]bool, n+1) 			// The length is n+1 coz we dont use the index 0 (num is from 1 to n)

	for i := 1; i <= n - 2; i++ {		
		var num int
		fmt.Scan(&num)
		exist[num] = true 				// Note that the default value for bool is false
	} 

	for i := 1; i <= n; i++ {
		if !exist[i] {					// If the current value of exist is false, print it
			fmt.Println(i)
		}
	}
}
