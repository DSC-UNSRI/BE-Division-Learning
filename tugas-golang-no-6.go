package main

import (
	"fmt"
)

type pair struct {
	first, second int
}

func backtrack(matches [][2]int, current []int, idx int, target []int) bool {
	if idx == len(matches) {
		for i := 0; i < len(current); i++ {
			if current[i] != target[i] {
				return false
			}
		}
		return true
	}

	i := matches[idx][0]
	j := matches[idx][1]

	// Tim i menang
	current[i] += 3
	if current[i] <= target[i] {
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[i] -= 3

	// Tim j menang
	current[j] += 3
	if current[j] <= target[j] {
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[j] -= 3

	// Seri
	current[i]++
	current[j]++
	if current[i] <= target[i] && current[j] <= target[j] {
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[i]--
	current[j]--

	return false
}

func main() {
	var T int
	fmt.Scan(&T)

	for testCase := 0; testCase < T; testCase++ {
		var N int
		fmt.Scan(&N)

		target := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&target[i])
		}

		var matches [][2]int
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				matches = append(matches, [2]int{i, j})
			}
		}

		current := make([]int, N)

		if backtrack(matches, current, 0, target) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
