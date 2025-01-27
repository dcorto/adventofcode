package main

import (
	"fmt"
	"time"
	"utils"
)

const day = 17

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

	containers := getContainers(lines)
	var combinations [][]int
	for i := 1; i <= len(containers); i++ {
		combinations = append(combinations, combinationsOfSum(containers, i, 150)...)
	}
	solution = len(combinations)
	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	containers := getContainers(lines)
	var combinations [][]int
	for i := 1; i <= len(containers); i++ {
		combinations = append(combinations, combinationsOfSum(containers, i, 150)...)
	}

	minLen := len(containers)
	for _, c := range combinations {
		if len(c) < minLen {
			minLen = len(c)
		}
	}

	for _, c := range combinations {
		if len(c) == minLen {
			solution++
		}
	}

	return solution
}

// getContainers returns a slice of integers from a slice of strings
func getContainers(lines []string) []int {
	var containers []int
	for _, line := range lines {
		containers = append(containers, utils.Atoi(line))
	}
	return containers
}

// combinationsOfSum returns all combinations of n elements from containers that sum up to target
func combinationsOfSum(containers []int, n, target int) [][]int {
	var result [][]int
	combine(containers, []int{}, n, target, &result)
	return result
}

// combine generates all combinations of n elements from containers that sum up to target
func combine(containers, current []int, n, target int, result *[][]int) {
	if target == 0 && len(current) == n {
		*result = append(*result, append([]int(nil), current...))
		return
	}
	if target < 0 || len(current) > n {
		return
	}
	for i, c := range containers {
		combine(containers[i+1:], append(current, c), n, target-c, result)
	}
}
