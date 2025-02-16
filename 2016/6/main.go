package main

import (
	"fmt"
	"math"
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

func solutionA() string {
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	var solution = make([]string, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		var freq = make(map[string]int)
		for j := 0; j < len(lines); j++ {
			freq[string(lines[j][i])]++
		}

		var max = 0
		var maxChar = ""
		for k, v := range freq {
			if v > max {
				max = v
				maxChar = k
			}
		}

		solution[i] = maxChar
	}

	return strings.Join(solution, "")
}

func solutionB() string {
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	var solution = make([]string, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		var freq = make(map[string]int)
		for j := 0; j < len(lines); j++ {
			freq[string(lines[j][i])]++
		}

		var min = math.MaxInt
		var minChar = ""
		for k, v := range freq {
			if v < min {
				min = v
				minChar = k
			}
		}

		solution[i] = minChar
	}

	return strings.Join(solution, "")
}
