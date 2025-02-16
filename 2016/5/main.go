package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func solutionA() string {
	var solution = ""

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for i := 0; i < 100000000; i++ {
		hash := md5Hash(input + fmt.Sprintf("%d", i))
		if hash[:5] == "00000" {
			solution = solution + string(hash[5])
			if len(solution) == 8 {
				break
			}
		}
	}

	return solution
}

func solutionB() string {
	var solution = make([]string, 8)

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	for i := 0; i < 100000000; i++ {
		hash := md5Hash(input + fmt.Sprintf("%d", i))
		if hash[:5] == "00000" {
			pos, err := strconv.Atoi(string(hash[5]))
			if err != nil || pos >= 8 || solution[pos] != "" {
				continue
			}

			solution[pos] = string(hash[6])
			if isFilled(solution) {
				break
			}
		}
	}

	return strings.Join(solution, "")
}

// isFilled checks if all elements in the array are filled
func isFilled(arr []string) bool {
	for _, v := range arr {
		if v == "" {
			return false
		}
	}
	return true
}

// md5Hash returns the md5 hash of the given data
func md5Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
