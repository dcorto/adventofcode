package main

import (
	"fmt"
	"sort"
	"time"
	"utils"
)

const day = 24

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

	weights := createWeightsFromLines(lines)
	solution = run(3, weights)
	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	weights := createWeightsFromLines(lines)
	solution = run(4, weights)
	return solution
}

// run divides the weights into groups and returns the product of the weights in the first group that has the same sum as the other groups
func run(groups int, weights []int) int {
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}

	if totalWeight%groups != 0 {
		fmt.Println(fmt.Sprintf("Cannot divide weights into %d groups with equal sum", groups))
		return 0
	}

	target := totalWeight / groups
	subsets := findSubsets(weights, target)
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		productI := 1
		for _, val := range subsets[i] {
			productI *= val
		}
		productJ := 1
		for _, val := range subsets[j] {
			productJ *= val
		}
		return productI < productJ
	})

	for _, subset := range subsets {
		remainingWeights := make(map[int]bool)
		for _, weight := range weights {
			remainingWeights[weight] = true
		}
		for _, weight := range subset {
			delete(remainingWeights, weight)
		}
		remainingSlice := make([]int, 0, len(remainingWeights))
		for weight := range remainingWeights {
			remainingSlice = append(remainingSlice, weight)
		}
		remainingSubsets := findSubsets(remainingSlice, target)
		if len(remainingSubsets) > 0 {
			product := 1
			for _, val := range subset {
				product *= val
			}
			return product
		}
	}

	fmt.Println("No valid subsets found")
	return 0
}

// createWeightsFromLines converts the lines of the input file into a slice of integers
func createWeightsFromLines(lines []string) []int {
	var weights []int
	for _, line := range lines {
		weight := utils.Atoi(line)
		weights = append(weights, weight)
	}
	return weights
}

// findSubsets returns all subsets of the weights that sum up to the target
func findSubsets(weights []int, target int) [][]int {
	var subsets [][]int
	var subset []int
	var helper func(int, int)
	helper = func(start, sum int) {
		if sum == target {
			subsets = append(subsets, append([]int(nil), subset...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(weights); i++ {
			subset = append(subset, weights[i])
			helper(i+1, sum+weights[i])
			subset = subset[:len(subset)-1]
		}
	}
	helper(0, 0)
	return subsets
}
