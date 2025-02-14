package main

import (
	"fmt"
	"math"
	"strings"
	"time"
	"utils"
)

const day = 1

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

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	instructions := strings.Split(input, ", ")

	var degrees = 0
	var x, y = 0, 0

	for _, instruction := range instructions {
		//fmt.Println(instruction)
		steps := utils.Atoi(instruction[1:])
		if instruction[0] == 'R' {
			degrees += 90
		} else {
			degrees -= 90
		}

		switch degrees % 360 {
		case 0:
			y += steps
			break
		case 90, -270:
			x += steps
			break
		case 180, -180:
			y -= steps
			break
		case 270, -90:
			x -= steps
			break
		}
	}

	solution = int(math.Abs(float64(x)) + math.Abs(float64(y)))

	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	instructions := strings.Split(input, ", ")

	var degrees = 0
	var x, y = 0, 0

	var visited = make(map[string]struct{})
	visited["0,0"] = struct{}{}
	for _, instruction := range instructions {
		steps := utils.Atoi(instruction[1:])
		if instruction[0] == 'R' {
			degrees += 90
		} else {
			degrees -= 90
		}

		for i := 0; i < steps; i++ {
			switch degrees % 360 {
			case 0:
				y++
			case 90, -270:
				x++
			case 180, -180:
				y--
			case 270, -90:
				x--
			}

			location := fmt.Sprintf("%d,%d", x, y)
			if _, ok := visited[location]; ok {
				return int(math.Abs(float64(x)) + math.Abs(float64(y)))
			}
			visited[location] = struct{}{}
		}
	}

	return solution
}
