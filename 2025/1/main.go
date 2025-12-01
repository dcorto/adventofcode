package main

import (
	"fmt"
	"strconv"
	"strings"

	"utils"
)

const day = 1

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	dial := 50 //dial start
	for _, line := range lines {
		distance, _ := strconv.Atoi(line[1:])
		distance = distance % 100 // when distance >= 100
		if strings.ToLower(string(line[0])) == "l" {
			dial = dial - distance
			if dial < 0 {
				dial = 100 + dial
			}
		} else {
			dial = dial + distance
			if dial > 99 {
				dial = dial - 100
			}
		}
		if dial == 0 {
			solution += 1
		}
	}

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	dial := 50 //dial start
	for _, line := range lines {
		v, _ := strconv.Atoi(line[1:])
		solution = solution + v/100
		v = v % 100
		if strings.ToLower(string(line[0])) == "l" {
			dial = dial - v
			if dial < 0 {
				dial = 100 + dial
				if dial+v != 100 {
					solution += 1
				}
			} else if dial == 0 {
				solution += 1
			}
		} else {
			dial = dial + v
			if dial > 99 {
				dial = dial - 100
				if dial-v != 0 {
					solution += 1
				}
			} else if dial == 0 {
				solution += 1
			}
		}
	}

	return solution
}
