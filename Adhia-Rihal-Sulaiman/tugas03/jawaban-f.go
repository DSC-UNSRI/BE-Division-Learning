package main

import "fmt"


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
	

	current[i] += 3
	if current[i] <= target[i] { 
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[i] -= 3
	

	current[j] += 3
	if current[j] <= target[j] {
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[j] -= 3
	

	current[i] += 1
	current[j] += 1
	if current[i] <= target[i] && current[j] <= target[j] { 
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[i] -= 1
	current[j] -= 1
	

	return false
}

func main() {
	var t int
	fmt.Scan(&t)
	

	for tc := 0; tc < t; tc++ {
		var n int

		fmt.Scan(&n)
		

		target := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&target[i])
		}
		

		var matches [][2]int
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				matches = append(matches, [2]int{i, j})
			}
		}
		

		current := make([]int, n)
		

		if backtrack(matches, current, 0, target) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
