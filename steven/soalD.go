package main

import (
	"fmt"
)

func main() {
	var A int;
	fmt.Scan(&A)

	bilanganAda := make([]bool, A+1)

	for i := 0; i < A-2; i++ {
		var data int;
		fmt.Scan(&data)
		
		bilanganAda[data] = true
	}
	
	bilanganHilang := []int{}
	for i := 1; i <= A; i++ {
		if !bilanganAda[i] {
			bilanganHilang = append(bilanganHilang, i)
		}
	}
	fmt.Println(bilanganHilang)
}