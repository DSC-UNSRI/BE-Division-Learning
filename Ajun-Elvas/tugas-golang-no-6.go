package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Fungsi untuk mengecek apakah tebakan skor memungkinkan
func fact(n int, nilai []int) bool {
	// Jumlah total skor harus kelipatan dari 3
	jumlahnilai := 0
	for _, score := range nilai {
		jumlahnilai += score
	}
	if jumlahnilai%3 != 0 {
		return false
	}

	// Sort skor dari terbesar ke terkecil
	sort.Sort(sort.Reverse(sort.IntSlice(nilai)))

	// Cek apakah mungkin mendistribusikan kemenangan secara adil
	for i := 0; i < n; i++ {
		if nilai[i] > (n - 1) { // Tim tidak bisa menang lebih dari jumlah lawannya
			return false
		}

		// Kurangi kemenangan dari tim-tim berikutnya
		for j := i + 1; j < i+1+nilai[i] && j < n; j++ {
			nilai[j]--
			if nilai[j] < 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Baca jumlah tebakan
	tStr, _ := reader.ReadString('\n')
	T, _ := strconv.Atoi(strings.TrimSpace(tStr))

	// Proses tiap tebakan
	for i := 0; i < T; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		// Ambil jumlah tim
		N, _ := strconv.Atoi(parts[0])
		nilai := make([]int, N)

		// Ambil skor-skor tim
		for j := 0; j < N; j++ {
			nilai[j], _ = strconv.Atoi(parts[j+1])
		}

		// Cek kemungkinan skor dan print hasilnya
		if fact(N, nilai) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
