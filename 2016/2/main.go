package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const (
	day = 2
)

// Keypad represents a keypad with buttons
type Keypad struct {
	buttons [][]string
}

// Position represents a position on the keypad
type Position struct {
	x, y int
}

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

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var kpWidth = 3
	var kpHeight = 3
	var kp = Keypad{
		buttons: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
	}

	code := make([]string, 0)
	p := Position{x: 1, y: 1} // Start at 5
	for _, line := range lines {
		for _, direction := range line {
			newP := move(p, string(direction))
			if isInside(newP, kp, kpWidth, kpHeight) {
				p = newP
			}
		}
		code = append(code, kp.buttons[p.x][p.y])
	}

	solution = strings.Join(code, "")
	return solution
}

func solutionB() string {
	var solution = ""

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var kpWidth = 5
	var kpHeight = 5
	var kp = Keypad{
		buttons: [][]string{
			{"", "", "1", "", ""},
			{"", "2", "3", "4", ""},
			{"5", "6", "7", "8", "9"},
			{"", "A", "B", "C", ""},
			{"", "", "D", "", ""},
		},
	}

	code := make([]string, 0)
	p := Position{x: 2, y: 0} // Start at 5
	for _, line := range lines {
		for _, direction := range line {
			newP := move(p, string(direction))
			if isInside(newP, kp, kpWidth, kpHeight) {
				p = newP
			}
		}
		code = append(code, kp.buttons[p.x][p.y])
	}

	solution = strings.Join(code, "")
	return solution
}

// move moves the position in the given direction
func move(p Position, direction string) Position {
	switch direction {
	case "U":
		p.x--
	case "D":
		p.x++
	case "L":
		p.y--
	case "R":
		p.y++
	}
	return p
}

// isInside checks if the position is inside the keypad and is not empty
func isInside(p Position, keypad Keypad, width, height int) bool {
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height && keypad.buttons[p.x][p.y] != ""
}
