package main

import "fmt"

func main() {
	var jumlahBilangan int
	var bilanganHilang []int

	fmt.Scan(&jumlahBilangan)

	bilanganAda := make(map[int]bool)
	for i := 0; i < jumlahBilangan-2; i++ {
		var bilangan int
		fmt.Scan(&bilangan)
		bilanganAda[bilangan] = true
	}
	for i := 1; i <= jumlahBilangan; i++ {
		if !bilanganAda[i] {
			bilanganHilang = append(bilanganHilang, i)
		}
	}

	fmt.Println(bilanganHilang[0])
	fmt.Println(bilanganHilang[1])
}
