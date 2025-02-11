package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"utils"
)

const day = 23

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

	instructions, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	registers := map[string]int{
		"a": 0,
		"b": 0,
	}

	run(instructions, registers)

	solution = registers["b"]
	return solution
}

func solutionB() int {
	var solution = 0

	instructions, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	registers := map[string]int{
		"a": 1, // Part 2 only change
		"b": 0,
	}

	run(instructions, registers)

	solution = registers["b"]
	return solution
}

func run(instructions []string, registers map[string]int) {
	i := 0
	for i < len(instructions) {
		parts := strings.SplitN(instructions[i], " ", 2)
		cmd := parts[0]
		param := parts[1]
		switch cmd {
		case "hlf":
			registers[param] = registers[param] / 2
			i++
		case "tpl":
			registers[param] *= 3
			i++
		case "inc":
			registers[param]++
			i++
		case "jmp":
			offset, _ := strconv.Atoi(param)
			i += offset
		case "jie":
			params := strings.Split(param, ", ")
			p1 := params[0]
			offset, _ := strconv.Atoi(params[1])
			if registers[p1]%2 == 0 {
				i += offset
			} else {
				i++
			}
		case "jio":
			params := strings.Split(param, ", ")
			p1 := params[0]
			offset, _ := strconv.Atoi(params[1])
			if registers[p1] == 1 {
				i += offset
			} else {
				i++
			}
		}
	}
}
