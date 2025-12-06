package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"utils"
)

const day = 5

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
		return 0
	}

	ranges := make([][2]int, 0)
	ingredients := make([]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				ranges = append(ranges, [2]int{start, end})
			}
			continue
		}
		i, _ := strconv.Atoi(line)
		ingredients = append(ingredients, i)
	}

	for _, i := range ingredients {
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				solution++
				break
			}
		}
	}

	return solution
}

func solutionB() int64 {
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	ranges := make([][2]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				ranges = append(ranges, [2]int{start, end})
			}
		}
	}

	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by their start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// Merge overlapping ranges
	merged := make([][2]int, 0)
	merged = append(merged, ranges[0])

	for i := 1; i < len(ranges); i++ {
		lastMerged := &merged[len(merged)-1]
		current := ranges[i]

		if current[0] <= (*lastMerged)[1] {
			if current[1] > (*lastMerged)[1] {
				(*lastMerged)[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	// Calculate the total size from the merged ranges
	var totalSize int64 = 0
	for _, r := range merged {
		totalSize += int64(r[1]-r[0]) + 1
	}

	return totalSize
}
