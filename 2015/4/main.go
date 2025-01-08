package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 4

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

	number := 0
	for {
		hash := getMD5Hash(fmt.Sprintf("%s%d", input, number))
		if strings.HasPrefix(hash, "00000") {
			break
		}
		number++
	}

	solution = number
	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	number := 0
	for {
		hash := getMD5Hash(fmt.Sprintf("%s%d", input, number))
		if strings.HasPrefix(hash, "000000") {
			break
		}
		number++
	}

	solution = number
	return solution
}

// getMD5Hash returns the MD5 hash of the input string
func getMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", hash)
}
