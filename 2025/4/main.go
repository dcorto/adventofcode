package main

import (
	"fmt"
	"strings"
	"time"

	"utils"
)

const day = 4

const roll = "@"
const removed = "x"

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

	board := make([][]string, 0)

	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	solution = findAccesibleRollsInBoard(board)

	return solution
}

func solutionB() int {
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	board := make([][]string, 0)
	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	totalRemoved := 0
	for {
		removedThisRound := findAndRemove(board)
		if removedThisRound == 0 {
			break
		}
		totalRemoved += removedThisRound
	}

	return totalRemoved
}

func findAndRemove(board [][]string) int {
	if len(board) == 0 {
		return 0
	}
	height := len(board)
	width := len(board[0])

	toRemove := make([][2]int, 0)

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if board[r][c] == roll {
				adjacents := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						nr, nc := r+dr, c+dc
						if isValidPosition(nr, nc, height, width) && board[nr][nc] == roll {
							adjacents++
						}
					}
				}

				if adjacents < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}
	}

	if len(toRemove) == 0 {
		return 0
	}

	for _, pos := range toRemove {
		board[pos[0]][pos[1]] = removed
	}

	return len(toRemove)
}

func findAccesibleRollsInBoard(board [][]string) int {
	if len(board) == 0 {
		return 0
	}
	height := len(board)
	width := len(board[0])

	found := 0

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if board[r][c] == roll {
				adjacents := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}

						nr, nc := r+dr, c+dc
						if isValidPosition(nr, nc, height, width) && board[nr][nc] == roll {
							adjacents++
						}
					}
				}

				if adjacents < 4 {
					found++
				}
			}
		}
	}

	return found
}

func isValidPosition(r, c, height, width int) bool {
	return r >= 0 && r < height && c >= 0 && c < width
}

// printBoard only for debug purposes
func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}
