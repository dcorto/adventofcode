package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 10

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

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	n := 40
	for i := 0; i < n; i++ {
		input = lookAndSay(input)
	}

	solution = len(input)

	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	n := 50
	for i := 0; i < n; i++ {
		input = lookAndSay(input)
	}

	solution = len(input)

	return solution
}

// lookAndSay returns the next number in the sequence
func lookAndSay(s string) string {
	var result strings.Builder
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result.WriteString(fmt.Sprintf("%d%c", count, s[i-1]))
			count = 1
		}
	}
	result.WriteString(fmt.Sprintf("%d%c", count, s[len(s)-1]))
	return result.String()
}
