package main

import "fmt"

func main() {
	var lengthInput int;
	fmt.Scan(&lengthInput);

	frekuensi := make(map[int]int);
	maxFrekuensi := 0;
	maxModus := 0;
	
	for i := 0; i <lengthInput; i++ {
		var data int;
		fmt.Scan(&data);
		frekuensi[data]++

		if frekuensi[data] > maxFrekuensi {
			maxFrekuensi = frekuensi[data]
			maxModus = data
		} else if frekuensi[data] == maxFrekuensi && data > maxModus {
			maxModus = data
		}
	} 
	fmt.Print(maxModus);
	
}
