package main

import (
	"fmt"
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

	slice := strings.Split(input, "")

	for _, s := range slice {
		if s == "(" {
			solution++
		} else {
			solution--
		}
	}

	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	slice := strings.Split(input, "")

	var position = 0
	for i, s := range slice {
		if s == "(" {
			position++
		} else {
			position--
		}

		if position == -1 {
			solution = i + 1
			break
		}
	}

	return solution
}
