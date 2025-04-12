package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	array := make(map[int]int)
	var angka, maxAngka, maxModus int
	
	for i := 0; i < n; i++{
		fmt.Scan(&angka)
		array[angka]++
		if array[angka] > maxAngka || (array[angka] == maxAngka && angka > maxModus){
			maxAngka = array[angka]
			maxModus = angka
		}
	}
	fmt.Print(maxModus)
}