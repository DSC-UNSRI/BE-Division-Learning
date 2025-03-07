package main

import "fmt"

type Match struct {
	teamA int
	teamB int
}

func dfs(matchIdx int, current []int, matches []Match, target []int, teamCount int, possible *bool) {
	if matchIdx == len(matches) {
		ok := true
		for i := 0; i < teamCount; i++ {
			if current[i] != target[i] {
				ok = false
				break
			}
		}
		if ok {
			*possible = true
		}
		return
	}

	if *possible {
		return
	}

	m := matches[matchIdx]

	current[m.teamA] += 3
	dfs(matchIdx+1, current, matches, target, teamCount, possible)
	current[m.teamA] -= 3
	if *possible {
		return
	}

	current[m.teamB] += 3
	dfs(matchIdx+1, current, matches, target, teamCount, possible)
	current[m.teamB] -= 3
	if *possible {
		return
	}

	current[m.teamA] += 1
	current[m.teamB] += 1
	dfs(matchIdx+1, current, matches, target, teamCount, possible)
	current[m.teamA] -= 1
	current[m.teamB] -= 1
}

func MainWorldCupGroup() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var n int
		fmt.Scan(&n)

		targetScores := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&targetScores[i])
		}

		var matches []Match
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				matches = append(matches, Match{teamA: i, teamB: j})
			}
		}

		currentScores := make([]int, n)
		foundSolution := false

		dfs(0, currentScores, matches, targetScores, n, &foundSolution)

		if foundSolution {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
