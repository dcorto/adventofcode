package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"utils"
)

const day = 2

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

	input, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	ranges := strings.Split(input[0], ",")
	for _, r := range ranges {
		val := strings.Split(r, "-")
		init, _ := strconv.Atoi(val[0])
		end, _ := strconv.Atoi(val[1])
		for i := init; i <= end; i++ {
			if checkIsNotValidProductID(strconv.Itoa(i)) {
				solution += i
			}
		}
	}

	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	ranges := strings.Split(input[0], ",")
	for _, r := range ranges {
		val := strings.Split(r, "-")
		init, _ := strconv.Atoi(val[0])
		end, _ := strconv.Atoi(val[1])
		for i := init; i <= end; i++ {
			if checkIsNotValidProductIDPart2(strconv.Itoa(i)) {
				solution += i
			}
		}
	}

	return solution
}

func checkIsNotValidProductID(s string) bool {
	return s[:len(s)/2] == s[len(s)/2:]
}
func checkIsNotValidProductIDPart2(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ { // i represents the length of the repeating substring
		if n%i == 0 { // Check if the string can be perfectly divided into substrings of length i
			firstSubstring := s[:i]
			allMatch := true
			for j := i; j < n; j += i {
				if s[j:j+i] != firstSubstring {
					allMatch = false
					break
				}
			}
			if allMatch {
				return true
			}
		}
	}
	return false
}
