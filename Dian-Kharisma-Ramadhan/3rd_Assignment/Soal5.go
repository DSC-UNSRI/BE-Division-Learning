package main

import "fmt"

func main() {
	var jumlahT int    
	var bilangan_N int

	fmt.Scan(&jumlahT)

	for i := 0; i < jumlahT; i++ {
		fmt.Scan(&bilangan_N)
		k4 := bilangan_N / 4
		jumlahKelipatan4 := 4 * k4 * (k4 + 1) / 2
		k7 := bilangan_N / 7
		jumlahKelipatan7 := 7 * k7 * (k7 + 1) / 2
		k28 := bilangan_N / 28
		jumlahKelipatan28 := 28 * k28 * (k28 + 1) / 2
		total := jumlahKelipatan4 + jumlahKelipatan7 - jumlahKelipatan28

		fmt.Println(total)
	}
}
