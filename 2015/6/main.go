package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 6

const (
	TurnOn  = "on"
	TurnOff = "off"
	Toggle  = "toggle"
)

var GridA [1000][1000]bool
var GridB [1000][1000]int

type Point struct {
	x int
	y int
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

func solutionA() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		action, p1, p2 := parseLine(line)
		switch action {
		case TurnOn:
			turnOn(p1.x, p1.y, p2.x, p2.y)
		case TurnOff:
			turnOff(p1.x, p1.y, p2.x, p2.y)
		case Toggle:
			toggle(p1.x, p1.y, p2.x, p2.y)
		}
	}

	solution = countLightsLit()

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		action, p1, p2 := parseLine(line)
		switch action {
		case TurnOn:
			increaseBrightness(p1.x, p1.y, p2.x, p2.y)
		case TurnOff:
			decreaseBrightness(p1.x, p1.y, p2.x, p2.y)
		case Toggle:
			doubleBrightness(p1.x, p1.y, p2.x, p2.y)
		}
	}

	solution = countLightsBrightness()
	return solution
}

// parseInputForPartOne parses the input
func parseLine(line string) (string, Point, Point) {
	var action string
	var p1, p2 Point

	if strings.Contains(line, "turn") {
		_, err := fmt.Sscanf(line, "turn %s %d,%d through %d,%d", &action, &p1.x, &p1.y, &p2.x, &p2.y)
		if err != nil {
			fmt.Println("Error:", err)
			return "", Point{}, Point{}
		}
	} else {
		_, err := fmt.Sscanf(line, "%s %d,%d through %d,%d", &action, &p1.x, &p1.y, &p2.x, &p2.y)
		if err != nil {
			fmt.Println("Error:", err)
			return "", Point{}, Point{}
		}
	}

	return action, p1, p2
}

// turnOn turns on the lights in the specified area
func turnOn(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			GridA[x][y] = true
		}
	}
}

// turnOff turns off the lights in the specified area
func turnOff(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			GridA[x][y] = false
		}
	}
}

// toggle toggles the lights in the specified area
func toggle(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			GridA[x][y] = !GridA[x][y]
		}
	}
}

// increaseBrightness increases the brightness of the lights in the specified area
func increaseBrightness(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			GridB[x][y]++
		}
	}
}

// decreaseBrightness decreases the brightness of the lights in the specified area
func decreaseBrightness(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			GridB[x][y]--
			if GridB[x][y] < 0 {
				GridB[x][y] = 0
			}
		}
	}
}

// doubleBrightness doubles the brightness of the lights in the specified area
func doubleBrightness(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			GridB[x][y] += 2
		}
	}
}

// countLightsLit counts the number of lights lit
func countLightsLit() int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if GridA[x][y] {
				count++
			}
		}
	}
	return count
}

// countLightsBrightness counts the total brightness of the lights
func countLightsBrightness() int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			count += GridB[x][y]
		}
	}
	return count
}
