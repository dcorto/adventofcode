package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"utils"
)

const day = 6

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

	for r, line := range lines {
		l := strings.Split(line, " ")
		if r < len(lines)-1 {
			numbers := make([]string, 0)
			for c := 0; c < len(l); c++ {
				_, err := strconv.Atoi(l[c])
				if err == nil {
					numbers = append(numbers, l[c])
				}
			}
			grid = append(grid, numbers)
			continue
		}

		op := make([]string, 0)
		for c := 0; c < len(l); c++ {
			if l[c] == "*" || l[c] == "+" {
				op = append(op, l[c])
			}
		}
		grid = append(grid, op)

	}

	height := len(grid)
	width := len(grid[0])
	operatorRow := grid[height-1]

	for c := 0; c < width; c++ {
		op := operatorRow[c]
		var partial int
		if op == "+" {
			for r := 0; r < height-1; r++ {
				num, _ := strconv.Atoi(grid[r][c])
				partial += num
			}
		} else if op == "*" {
			partial = 1 // Start with 1 for multiplication
			for r := 0; r < height-1; r++ {
				num, _ := strconv.Atoi(grid[r][c])
				partial *= num
			}
		}
		solution += partial
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

	height := len(grid)
	width := len(grid[0])

	var numbers []int
	for c := width - 1; c >= 0; c-- {
		digits := ""
		for r := 0; r < height; r++ {
			if r == height-1 { //operation row
				_, err := strconv.Atoi(grid[r][c])
				if err != nil {
					if grid[r][c] == "+" {
						n, _ := strconv.Atoi(digits)
						numbers = append(numbers, n)
						partial := 0
						for _, n := range numbers {
							partial += n
						}
						solution += partial
						numbers = make([]int, 0)
						digits = ""
						continue
					}

					if grid[r][c] == "*" {
						n, _ := strconv.Atoi(digits)
						numbers = append(numbers, n)
						partial := 1
						for _, n := range numbers {
							partial *= n
						}
						solution += partial
						numbers = make([]int, 0)
						digits = ""
						continue
					}

					if grid[r][c] == " " && digits != "" {
						n, _ := strconv.Atoi(digits)
						numbers = append(numbers, n)
					}
				}
			}
			_, err := strconv.Atoi(grid[r][c])
			if err != nil {
				continue
			}
			digits += grid[r][c]
		}
	}
	return solution
}

// printBoard only for debug purposes
func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}
