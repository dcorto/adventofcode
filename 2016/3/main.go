package main

import (
	"fmt"
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

func solutionA() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var a, b, c int
	for _, line := range lines {
		fmt.Sscanf(line, "%d%d%d", &a, &b, &c)
		if valid(a, b, c) {
			solution++
		}
	}

	return solution
}

func solutionB() int {
	var solution = 0

	data, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var a1, b1, c1, a2, b2, c2, a3, b3, c3 int
	for i := 0; i < len(data); i += 3 {
		fmt.Sscanf(data[i], "%d%d%d", &a1, &a2, &a3)
		fmt.Sscanf(data[i+1], "%d%d%d", &b1, &b2, &b3)
		fmt.Sscanf(data[i+2], "%d%d%d", &c1, &c2, &c3)
		if valid(a1, b1, c1) {
			solution++
		}
		if valid(a2, b2, c2) {
			solution++
		}
		if valid(a3, b3, c3) {
			solution++
		}

	}

	return solution
}

// valid checks if the given sides can form a valid triangle
func valid(a int, b int, c int) bool {
	return a+b > c && a+c > b && c+b > a
}
