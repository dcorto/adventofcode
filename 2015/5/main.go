package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 5

func main() {
	fmt.Println("Solution for Day", day)

	startTimeA := time.Now()
	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA, "(Time:", time.Since(startTimeA), ")")

	startTimeB := time.Now()
	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB, "(Time:", time.Since(startTimeB), ")")
}

func solutionA() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	niceStrings := 0
	for _, line := range lines {
		if isNiceString(line) {
			niceStrings++
		}
	}

	solution = niceStrings

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	niceStrings := 0
	for _, line := range lines {
		if isBetterNiceString(line) {
			niceStrings++
		}
	}

	solution = niceStrings

	return solution
}

// isNiceString checks if a string is nice according to the rules
func isNiceString(s string) bool {
	vowels := "aeiou"
	forbidden := []string{"ab", "cd", "pq", "xy"}
	vowelCount := 0
	hasDouble := false

	for i := 0; i < len(s); i++ {
		// Count vowels
		if strings.ContainsRune(vowels, rune(s[i])) {
			vowelCount++
		}

		// Check double letters
		if i > 0 && s[i] == s[i-1] {
			hasDouble = true
		}

		// Check forbidden pairs
		if i > 0 {
			pair := s[i-1 : i+1]
			for _, f := range forbidden {
				if pair == f {
					return false
				}
			}
		}
	}

	return vowelCount >= 3 && hasDouble
}

// isBetterNiceString checks if a string is nice according to the new rules
func isBetterNiceString(s string) bool {
	hasDouble := false
	hasGap := false

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Count(s, pair) >= 2 {
			hasDouble = true
			break
		}
	}

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			hasGap = true
			break
		}
	}

	return hasDouble && hasGap
}
