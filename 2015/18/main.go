package main

import (
	"fmt"
	"time"
	"utils"
)

const day = 18

const (
	On         = "#"
	Off        = "."
	widthMax   = 100
	heightMax  = 100
	iterations = 100
)

type Grid [widthMax][heightMax]string

var grid Grid

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

	createGrid(lines, &grid)

	for i := 0; i < iterations; i++ {

		currentGrid := copyGrid(&grid)

		for y := 0; y < heightMax; y++ {
			for x := 0; x < widthMax; x++ {

				// Count neighbors
				neighbors := countNeighbors(x, y, &currentGrid)

				// Apply rules
				if grid[x][y] == On && neighbors != 2 && neighbors != 3 {
					grid[x][y] = Off
					continue
				}

				if grid[x][y] == Off && neighbors == 3 {
					grid[x][y] = On
					continue
				}
			}
		}
	}

	solution = countLightsOn(&grid)

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	createGrid(lines, &grid)

	setCornersOn(&grid)
	for i := 0; i < iterations; i++ {
		currentGrid := copyGrid(&grid)

		for y := 0; y < heightMax; y++ {
			for x := 0; x < widthMax; x++ {

				// Count neighbors
				neighbors := countNeighbors(x, y, &currentGrid)

				// Apply rules
				if grid[x][y] == On && neighbors != 2 && neighbors != 3 {
					grid[x][y] = Off
					continue
				}

				if grid[x][y] == Off && neighbors == 3 {
					grid[x][y] = On
					continue
				}
			}
		}
		setCornersOn(&grid)
	}
	solution = countLightsOn(&grid)

	return solution
}

// setCornersOn sets the lights corners of the grid to On
func setCornersOn(grid *Grid) {
	grid[0][0] = On
	grid[0][heightMax-1] = On
	grid[widthMax-1][0] = On
	grid[widthMax-1][heightMax-1] = On
}

// countLightsOn counts the number of lights on
func countLightsOn(grid *Grid) int {
	var count = 0
	for y := 0; y < heightMax; y++ {
		for x := 0; x < widthMax; x++ {
			if grid[x][y] == On {
				count++
			}
		}
	}
	return count
}

// copyGrid copies the grid
func copyGrid(grid *Grid) Grid {
	var newGrid Grid
	for y := 0; y < heightMax; y++ {
		for x := 0; x < widthMax; x++ {
			newGrid[x][y] = grid[x][y]
		}
	}
	return newGrid
}

// countNeighbors counts the number of neighbors
func countNeighbors(x, y int, grid *Grid) int {
	neighbors := 0
	for y2 := y - 1; y2 <= y+1; y2++ {
		for x2 := x - 1; x2 <= x+1; x2++ {
			if x2 < 0 || x2 >= widthMax || y2 < 0 || y2 >= heightMax {
				continue
			}
			if x2 == x && y2 == y {
				continue
			}
			if grid[x2][y2] == On {
				neighbors++
			}
		}
	}
	return neighbors
}

// createGrid creates a grid from the input
func createGrid(lines []string, grid *Grid) {
	for y, line := range lines {
		for x, char := range line {
			grid[x][y] = string(char)
		}
	}
}
