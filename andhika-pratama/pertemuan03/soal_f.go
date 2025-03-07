package main

import "fmt"

// backtrack recursively assigns outcomes to each match and checks if the
// resulting points for each team can match the target configuration.
// - matches: list of pairs (i, j) representing each match between team i and team j.
// - current: current accumulated points for each team.
// - idx: the current match index we are assigning an outcome to.
// - target: the target final scores for each team.
func backtrack(matches [][2]int, current []int, idx int, target []int) bool {
	// If all matches have been processed, check if we reached the target scores.
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

	// Outcome 1: Team i wins (i gets 3 points, j gets 0)
	current[i] += 3
	if current[i] <= target[i] { // prune if team i exceeds its target
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[i] -= 3

	// Outcome 2: Team j wins (j gets 3 points, i gets 0)
	current[j] += 3
	if current[j] <= target[j] { // prune if team j exceeds its target
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[j] -= 3

	// Outcome 3: Draw (both teams get 1 point)
	current[i] += 1
	current[j] += 1
	if current[i] <= target[i] && current[j] <= target[j] { // prune if any team exceeds its target
		if backtrack(matches, current, idx+1, target) {
			return true
		}
	}
	current[i] -= 1
	current[j] -= 1

	// If no outcome leads to a valid configuration, return false.
	return false
}

func main() {
	var t int
	fmt.Scan(&t)

	// Process each test case.
	for tc := 0; tc < t; tc++ {
		var n int
		// Read the number of teams in the group.
		fmt.Scan(&n)

		// Read the target final scores for each team.
		target := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&target[i])
		}

		// Generate a list of all matches.
		// Each pair (i, j) with i < j represents a match between team i and team j.
		var matches [][2]int
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				matches = append(matches, [2]int{i, j})
			}
		}

		// Create a slice to keep track of the current points for each team.
		current := make([]int, n)

		// Use backtracking to check if the target configuration is possible.
		if backtrack(matches, current, 0, target) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}