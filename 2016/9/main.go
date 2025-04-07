package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"utils"
)

const day = 9

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

	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	decompressedLength := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			markerEnd := strings.IndexByte(data[i:], ')')
			marker := data[i+1 : i+markerEnd]
			parts := strings.Split(marker, "x")
			chars, _ := strconv.Atoi(parts[0])
			repeat, _ := strconv.Atoi(parts[1])

			decompressedLength += chars * repeat
			i += markerEnd + chars
		} else {
			decompressedLength++
		}
	}

	solution = decompressedLength

	return solution
}

func solutionB() int {
	var solution = 0
	return solution
}
