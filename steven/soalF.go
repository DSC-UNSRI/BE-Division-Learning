package main

import (
    "fmt"
)

type Match struct {
    teamA, teamB int
}

func simulateMatches(current int, tempScores []int, matches []Match, totalMatches int, finalScores []int, teams int, result *bool) {
    if current >= totalMatches {
        match := true
        for i := 0; i < teams; i++ {
            if tempScores[i] != finalScores[i] {
                match = false
                break
            }
        }
        if match {
            *result = true
        }
        return
    }

    A, B := matches[current].teamA, matches[current].teamB
    
    newScores := make([]int, teams)
    copy(newScores, tempScores)
    newScores[A] += 3
    simulateMatches(current+1, newScores, matches, totalMatches, finalScores, teams, result)
    if *result {
        return
    }

    copy(newScores, tempScores)
    newScores[B] += 3
    simulateMatches(current+1, newScores, matches, totalMatches, finalScores, teams, result)
    if *result {
        return
    }
    
    copy(newScores, tempScores)
    newScores[A]++
    newScores[B]++
    simulateMatches(current+1, newScores, matches, totalMatches, finalScores, teams, result)
    if *result {
        return
    }
}

func main() {
    var testCases int
    fmt.Scan(&testCases)

    for t := 0; t < testCases; t++ {
        var teams int
        fmt.Scan(&teams)

        finalScores := make([]int, teams)
        for i := 0; i < teams; i++ {
            fmt.Scan(&finalScores[i])
        }

        totalMatches := (teams * (teams - 1)) / 2
        matches := make([]Match, 0, totalMatches)

        for i := 0; i < teams; i++ {
            for j := i + 1; j < teams; j++ {
                matches = append(matches, Match{i, j})
            }
        }

        currentScores := make([]int, teams)
        possible := false
        simulateMatches(0, currentScores, matches, totalMatches, finalScores, teams, &possible)

        if possible {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}