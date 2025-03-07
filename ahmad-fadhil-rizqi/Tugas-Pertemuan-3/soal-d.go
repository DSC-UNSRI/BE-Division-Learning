package main

import (
	"fmt"
)

func main() {
	//deklarasi variabel
	var jumlahBilangan int
	var bilanganHilang []int

	//input jumlah bilangan
	fmt.Scan(&jumlahBilangan)

	//membuat map untuk menyimpan bilangan yang ada
	bilanganYangAda := make(map[int]bool)
	for i := 0; i < jumlahBilangan-2; i++ {
		var bilangan int
		fmt.Scan(&bilangan)
		bilanganYangAda[bilangan] = true
	}

	//mencari bilangan yang hilang
	for i := 1; i <= jumlahBilangan; i++ {
		if !bilanganYangAda[i] {
			bilanganHilang = append(bilanganHilang, i)
		}
	}

	//menampilkan bilangan yang hilang
	fmt.Println(bilanganHilang[0])
	fmt.Println(bilanganHilang[1])
}