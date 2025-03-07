package main

import "fmt"

func bruteForce(pertandinganKe int, skorSementara []int, P []pair, banyakPertandingan int, S []int, N int, Answer *bool) {
	if pertandinganKe >= banyakPertandingan {
		isArraySame := true
		for i := 0; i < N; i++ {
			if skorSementara[i] != S[i] {
				isArraySame = false
			}
		}

		if isArraySame {
			*Answer = true
		}
		return
	}

	A := P[pertandinganKe].first
	B := P[pertandinganKe].second

	skorBaru := make([]int, 5) // Ukuran tetap

	// tim A menang
	copy(skorBaru, skorSementara)
	skorBaru[A] += 3
	bruteForce(pertandinganKe+1, skorBaru, P, banyakPertandingan, S, N, Answer)
	if *Answer {
		return
	}

	// tim B menang
	copy(skorBaru, skorSementara)
	skorBaru[B] += 3
	bruteForce(pertandinganKe+1, skorBaru, P, banyakPertandingan, S, N, Answer)
	if *Answer {
		return
	}

	// seri
	copy(skorBaru, skorSementara)
	skorBaru[A] += 1
	skorBaru[B] += 1
	bruteForce(pertandinganKe+1, skorBaru, P, banyakPertandingan, S, N, Answer)
	if *Answer {
		return
	}
}

type pair struct {
	first  int
	second int
}

func main() {
	var T int
	fmt.Scan(&T) // Baca jumlah kasus uji

	for testCase := 1; testCase <= T; testCase++ {
		var N int
		fmt.Scan(&N)

		S := make([]int, 5) // Ukuran tetap
		for i := 0; i < N; i++ {
			fmt.Scan(&S[i])
		}

		banyakPertandingan := 0
		P := make([]pair, 15) // Ukuran tetap (cukup besar)
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				P[banyakPertandingan] = pair{i, j}
				banyakPertandingan++
			}
		}

		C := make([]int, 5) // Ukuran tetap
		Answer := false     // Inisialisasi

		bruteForce(0, C, P, banyakPertandingan, S, N, &Answer) // Pass by reference

		if Answer {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}