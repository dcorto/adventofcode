package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 25

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

	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	slit := strings.Split(data, ",")
	row := utils.Atoi(slit[0])
	col := utils.Atoi(slit[1])

	// initial code
	code := 20151125

	// constants
	const multiplier = 252533
	const mod = 33554393

	// Calculate the position in the sequence
	position := calculatePosition(row, col)

	// Generate the code in the given position
	for i := 1; i < position; i++ {
		code = (code * multiplier) % mod
	}
	solution = code
	return solution
}

func solutionB() int {
	var solution = 0
	return solution
}

// calculatePosition calculate the position in the cell (row, col)
func calculatePosition(row, col int) int {
	// the position in the sequence is calculated by adding the numbers of the diagonal
	diagonal := row + col - 1
	return (diagonal*(diagonal-1))/2 + col
}
