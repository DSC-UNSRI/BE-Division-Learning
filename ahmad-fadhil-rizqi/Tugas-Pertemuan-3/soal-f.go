package main

import (
	"fmt"
)

func main() {
	var jumlahTebakan int
	fmt.Scan(&jumlahTebakan)

	// Slice untuk menyimpan semua input
	tebakan := make([][]int, jumlahTebakan)

	// Membaca semua input
	for i := 0; i < jumlahTebakan; i++ {
		var jumlahTim int
		fmt.Scan(&jumlahTim)

		skorTarget := make([]int, jumlahTim)
		for j := 0; j < jumlahTim; j++ {
			fmt.Scan(&skorTarget[j])
		}

		tebakan[i] = skorTarget
	}

	// Fungsi backtracking untuk memeriksa kemungkinan konfigurasi skor
	var backtrack func(pertandingan [][2]int, skorSekarang []int, indeks int, skorTarget []int) bool
	backtrack = func(pertandingan [][2]int, skorSekarang []int, indeks int, skorTarget []int) bool {
		// Jika semua pertandingan telah diproses, periksa apakah skor sesuai target
		if indeks == len(pertandingan) {
			for i := 0; i < len(skorSekarang); i++ {
				if skorSekarang[i] != skorTarget[i] {
					return false
				}
			}
			return true
		}

		timA := pertandingan[indeks][0]
		timB := pertandingan[indeks][1]

		// Outcome 1: Tim A menang (Tim A dapat 3 poin, Tim B dapat 0)
		skorSekarang[timA] += 3
		if skorSekarang[timA] <= skorTarget[timA] && backtrack(pertandingan, skorSekarang, 
			indeks+1, skorTarget) {
			return true
		}
		skorSekarang[timA] -= 3

		// Outcome 2: Tim B menang (Tim B dapat 3 poin, Tim A dapat 0)
		skorSekarang[timB] += 3
		if skorSekarang[timB] <= skorTarget[timB] && backtrack(pertandingan, skorSekarang, 
			indeks+1, skorTarget) {
			return true
		}
		skorSekarang[timB] -= 3

		// Outcome 3: Seri (Tim A dan Tim B masing-masing dapat 1 poin)
		skorSekarang[timA] += 1
		skorSekarang[timB] += 1
		if skorSekarang[timA] <= skorTarget[timA] && skorSekarang[timB] <= skorTarget[timB] && backtrack(pertandingan, 
			skorSekarang, indeks+1, skorTarget) {
			return true
		}
		skorSekarang[timA] -= 1
		skorSekarang[timB] -= 1

		// Jika tidak ada outcome yang valid, return false
		return false
	}

	// Memproses dan menampilkan output setelah semua input dimasukkan
	for _, skorTarget := range tebakan {
		jumlahTim := len(skorTarget)
		var pertandingan [][2]int

		// Generate semua pertandingan
		for i := 0; i < jumlahTim; i++ {
			for j := i + 1; j < jumlahTim; j++ {
				pertandingan = append(pertandingan, [2]int{i, j})
			}
		}

		// Inisialisasi skor sekarang
		skorSekarang := make([]int, jumlahTim)

		// Gunakan backtracking untuk memeriksa kemungkinan konfigurasi skor
		if backtrack(pertandingan, skorSekarang, 0, skorTarget) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}