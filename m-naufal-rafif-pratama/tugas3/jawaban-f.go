package main

import "fmt"

func bruteForce(matchIndex int, currentScores []int, matches []pair, totalMatches int, targetScores []int, numTeams int, answer *bool) {
	if matchIndex >= totalMatches {
		isArraySame := true
		for i := 0; i < numTeams; i++ {
			if currentScores[i] != targetScores[i] {
				isArraySame = false
			}
		}

		if isArraySame {
			*answer = true
		}
		return
	}

	teamA := matches[matchIndex].first
	teamB := matches[matchIndex].second

	newScores := make([]int, 5)

	copy(newScores, currentScores)
	newScores[teamA] += 3
	bruteForce(matchIndex+1, newScores, matches, totalMatches, targetScores, numTeams, answer)
	if *answer {
		return
	}

	copy(newScores, currentScores)
	newScores[teamB] += 3
	bruteForce(matchIndex+1, newScores, matches, totalMatches, targetScores, numTeams, answer)
	if *answer {
		return
	}

	copy(newScores, currentScores)
	newScores[teamA] += 1
	newScores[teamB] += 1
	bruteForce(matchIndex+1, newScores, matches, totalMatches, targetScores, numTeams, answer)
	if *answer {
		return
	}
}

type pair struct {
	first  int
	second int
}

func main() {
	var T int
	fmt.Scan(&T)
	for testCase := 1; testCase <= T; testCase++ {
		var N int
		fmt.Scan(&N)

		targetScores := make([]int, 5)
		for i := 0; i < N; i++ {
			fmt.Scan(&targetScores[i])
		}

		totalMatches := 0
		matches := make([]pair, 15)
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				matches[totalMatches] = pair{i, j}
				totalMatches++
			}
		}

		initialScores := make([]int, 5)
		answer := false

		bruteForce(0, initialScores, matches, totalMatches, targetScores, N, &answer)

		if answer {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}