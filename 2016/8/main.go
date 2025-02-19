package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const (
	day          = 8
	width        = 50
	height       = 6
	draw         = "rect"
	rotateRow    = "rotate row"
	rotateColumn = "rotate column"
)

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

	screen := createScreen()

	for _, line := range lines {
		action, a, b := parseLine(line)
		switch action {
		case draw:
			for i := 0; i < b; i++ {
				for j := 0; j < a; j++ {
					screen[i][j] = true
				}
			}
			break
		case rotateRow:
			for i := 0; i < b; i++ {
				var last = screen[a][width-1]
				for j := width - 1; j > 0; j-- {
					screen[a][j] = screen[a][j-1]
				}
				screen[a][0] = last
			}
		case rotateColumn:
			for i := 0; i < b; i++ {
				var last = screen[height-1][a]
				for j := height - 1; j > 0; j-- {
					screen[j][a] = screen[j-1][a]
				}
				screen[0][a] = last
			}
		}
	}

	solution = countPixelsOn(screen)
	return solution
}

func solutionB() int {
	var solution = 0
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	screen := createScreen()

	for _, line := range lines {
		action, a, b := parseLine(line)
		switch action {
		case draw:
			for i := 0; i < b; i++ {
				for j := 0; j < a; j++ {
					screen[i][j] = true
				}
			}
			break
		case rotateRow:
			for i := 0; i < b; i++ {
				var last = screen[a][width-1]
				for j := width - 1; j > 0; j-- {
					screen[a][j] = screen[a][j-1]
				}
				screen[a][0] = last
			}
		case rotateColumn:
			for i := 0; i < b; i++ {
				var last = screen[height-1][a]
				for j := height - 1; j > 0; j-- {
					screen[j][a] = screen[j-1][a]
				}
				screen[0][a] = last
			}
		}
	}

	printScreen(screen)

	solution = countPixelsOn(screen)
	return solution
}

// createScreen creates a screen with the given dimensions
func createScreen() [][]bool {
	screen := make([][]bool, height)
	for i := range screen {
		screen[i] = make([]bool, width)
	}
	return screen
}

// printScreen prints the screen
func printScreen(screen [][]bool) {
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// countPixelsOn counts the number of pixels that are on
func countPixelsOn(screen [][]bool) int {
	var count = 0
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				count++
			}
		}
	}
	return count
}

// parseLine parses a line and returns the action and the values
func parseLine(line string) (string, int, int) {
	var action string
	var a, b int
	p := strings.Split(line, " ")
	if p[0] == draw {
		action = draw
		fmt.Sscanf(p[1], "%dx%d", &a, &b)
		return action, a, b
	}

	if p[1] == "row" {
		action = rotateRow
		fmt.Sscanf(strings.Join(p[2:], " "), "y=%d by %d", &a, &b)
		return action, a, b
	}
	if p[1] == "column" {
		action = rotateColumn
		fmt.Sscanf(strings.Join(p[2:], " "), "x=%d by %d", &a, &b)
		return action, a, b
	}
	return action, a, b
}
