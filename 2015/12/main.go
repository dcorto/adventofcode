package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"utils"
)

const day = 12

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

	regexp.MustCompile(`[\-0-9]+`).ReplaceAllStringFunc(input, func(match string) string {
		i, _ := strconv.Atoi(match)
		solution += i
		return match
	})

	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var f interface{}
	err = json.Unmarshal([]byte(input), &f)
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	solution = int(rec(f))

	return solution
}

// rec recursively sums all numbers in the given JSON object, except for those in objects containing the value "red"
func rec(f interface{}) (output float64) {
outer:
	switch fv := f.(type) {
	case []interface{}:
		for _, val := range fv {
			output += rec(val)
		}
	case float64:
		output += fv
	case map[string]interface{}:
		for _, val := range fv {
			if val == "red" {
				break outer
			}
		}
		for _, val := range fv {
			output += rec(val)
		}
	}

	return output
}
