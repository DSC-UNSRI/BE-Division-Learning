package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	counts := make(map[int]int) // Menyimpan frekuensi kemunculan setiap bilangan
	maxCount := 0               // Menyimpan frekuensi kemunculan tertinggi
	modusTerbesar := 0          // Menyimpan modus terbesar

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		counts[num]++ // Menambah frekuensi kemunculan bilangan
		
		if counts[num] > maxCount {
			maxCount = counts[num]
			modusTerbesar = num
		} else if counts[num] == maxCount && num > modusTerbesar {
			modusTerbesar = num // Jika frekuensi sama, pilih bilangan yang lebih besar
		}
	}

	fmt.Println(modusTerbesar)
}