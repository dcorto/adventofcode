package main

import (
	"fmt"
	"math"
	"time"
	"utils"
)

const day = 20

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

	minimum := utils.Atoi(data)

	var n = 1
	for {
		if sumOfPresentsPart1(n) >= minimum {
			break
		}
		n++

	}
	solution = n

	return solution
}

func sumOfPresentsPart1(n int) int {
	sum := 0
	d := int(math.Sqrt(float64(n))) + 1
	for i := 1; i <= d; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i { // Avoid adding the divisor twice if i * i == n
				sum += n / i
			}
		}
	}
	return sum * 10
}

func sumOfPresentsPart2(n int) int {
	sum := 0
	d := int(math.Sqrt(float64(n))) + 1
	for i := 1; i <= d; i++ {
		if n%i == 0 {
			if i <= 50 {
				sum += n / i
			}
			if n/i <= 50 {
				sum += i
			}
		}
	}
	return sum * 11
}

func solutionB() int {
	var solution = 0

	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	minimum := utils.Atoi(data)

	var n = 1
	for {
		if sumOfPresentsPart2(n) >= minimum {
			break
		}
		n++

	}
	solution = n

	return solution
}
