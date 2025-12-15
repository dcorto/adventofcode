package main

import (
	"fmt"
	"strings"
	"time"

	"utils"
)

const day = 7

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
		return 0
	}

	grid := make([][]string, 0)
	for _, line := range lines {
		l := strings.Split(line, "")
		grid = append(grid, l)
	}

	//printBoard(grid)

	height := len(grid)
	width := len(grid[0])

	for r := 1; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r-1][c] == "S" || grid[r-1][c] == "|" {
				if grid[r][c] == "^" {
					solution++
					grid[r][c+1] = "|"
					grid[r][c-1] = "|"
				} else {
					grid[r][c] = "|"
				}
			}
		}
		//printBoard(grid)
	}

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	grid := make([][]string, 0)
	for _, line := range lines {
		l := strings.Split(line, "")
		grid = append(grid, l)
	}

	// Find the starting column of 'S' in the first row.
	startCol := -1
	if len(grid) > 0 {
		for c, char := range grid[0] {
			if char == "S" {
				startCol = c
				break
			}
		}
	}

	if startCol == -1 {
		return 0
	}

	//printBoard(grid)

	cache := make(map[struct{ r, c int }]int)
	solution = recursive(grid, 0, startCol, cache)

	return solution
}

// recursive calculates the number of paths from a given (r, c) to the bottom of the grid.
func recursive(grid [][]string, r, c int, cache map[struct{ r, c int }]int) int {
	if c < 0 || c >= len(grid[0]) {
		return 0
	}

	if r >= len(grid) {
		return 1
	}

	if v, ok := cache[struct{ r, c int }{r, c}]; ok {
		return v
	}

	var result int
	if grid[r][c] == "^" {
		result = recursive(grid, r+1, c-1, cache) + recursive(grid, r+1, c+1, cache)
	} else {
		result = recursive(grid, r+1, c, cache)
	}

	cache[struct{ r, c int }{r, c}] = result
	return result
}

// printBoard only for debug purposes
func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}
