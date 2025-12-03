package main

import (
	"fmt"
	"strconv"
	"time"

	"utils"
)

const day = 3

func main() {
	fmt.Println("Solution for Day", day)

	startTimeA := time.Now()
	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA, "(Time:", time.Since(startTimeA), ")")

	startTimeB := time.Now()
	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB, "(Time:", time.Since(startTimeB), ")")
}

func solutionA() int64 {
	var solution int64 = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		joltage := findJoltageInLine(line, 2)
		solution += joltage
	}

	return solution
}

func solutionB() int64 {
	var solution int64 = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		joltage := findJoltageInLine(line, 12)
		solution += joltage
	}

	return solution
}

func findJoltageInLine(line string, digits int) int64 {
	k := digits
	if len(line) < k {
		num, _ := strconv.ParseInt(line, 10, 64)
		return num
	}

	removals := len(line) - k
	stack := make([]byte, 0, len(line))

	for i := 0; i < len(line); i++ {
		digit := line[i]
		for len(stack) > 0 && digit > stack[len(stack)-1] && removals > 0 {
			stack = stack[:len(stack)-1]
			removals--
		}
		stack = append(stack, digit)
	}

	if len(stack) > k {
		stack = stack[:k]
	}

	numStr := string(stack)
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
