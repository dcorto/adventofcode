package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 11

func main() {
	fmt.Println("Solution for Day", day)

	startTimeA := time.Now()
	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA, "(Time:", time.Since(startTimeA), ")")

	startTimeB := time.Now()
	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB, "(Time:", time.Since(startTimeB), ")")
}

func solutionA() string {
	var solution = ""

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	currentPassword := input
	solution = findNextPassword(currentPassword)

	return solution
}

func solutionB() string {
	var solution = ""

	currentPassword := "hxbxxyzz" // Solution A
	solution = findNextPassword(currentPassword)

	return solution
}

// findNextPassword returns the next valid password after the given one
func findNextPassword(password string) string {
	for {
		password = incrementPassword(password)
		if isValidPassword(password) {
			return password
		}
	}
}

// incrementPassword increments the given password by one
func incrementPassword(password string) string {
	runes := []rune(password)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
		} else {
			runes[i]++
			break
		}
	}
	return string(runes)
}

// isValidPassword returns true if the given password is valid
func isValidPassword(password string) bool {
	return hasIncreasingStraight(password) && !containsInvalidChars(password) && hasTwoNonOverlappingPairs(password)
}

// hasIncreasingStraight returns true if the given password contains an increasing straight of at least three letters
func hasIncreasingStraight(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i+1] == password[i]+1 && password[i+2] == password[i]+2 {
			return true
		}
	}
	return false
}

// containsInvalidChars returns true if the given password contains the letters 'i', 'o', or 'l'
func containsInvalidChars(password string) bool {
	return strings.ContainsAny(password, "iol")
}

// hasTwoNonOverlappingPairs returns true if the given password contains at least two non-overlapping pairs of letters
func hasTwoNonOverlappingPairs(password string) bool {
	pairCount := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairCount++
			i++ // Skip the next character to ensure pairs are non-overlapping
		}
	}
	return pairCount >= 2
}
