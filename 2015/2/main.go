package main

import (
	"fmt"
	"sort"
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

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		l, w, h := getDimensionsFromLine(line)
		area := getArea(l, w, h)
		smallestSide := getAreaOfSmallestSide(l, w, h)

		solution += area + smallestSide
	}

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		l, w, h := getDimensionsFromLine(line)

		ribbonToWrap := getRibbonToWrap(l, w, h)
		ribbonForBow := getRibbonForBow(l, w, h)

		solution += ribbonToWrap + ribbonForBow
	}

	return solution
}

// getDimensionsFromLine returns the dimensions of a box from a string
func getDimensionsFromLine(line string) (int, int, int) {
	slice := strings.Split(line, "x")
	l, _ := strconv.Atoi(slice[0])
	w, _ := strconv.Atoi(slice[1])
	h, _ := strconv.Atoi(slice[2])
	return l, w, h
}

// getAreaOfSmallestSide returns the area of the smallest side of a box
func getAreaOfSmallestSide(l, w, h int) int {
	sides := []int{l, w, h}
	sort.Ints(sides)
	return sides[0] * sides[1]
}

// getArea returns the area of a box
func getArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

// getRibbonToWrap returns the amount of ribbon needed to wrap a box
func getRibbonToWrap(l, w, h int) int {
	sides := []int{l, w, h}
	sort.Ints(sides)
	return sides[0] + sides[0] + sides[1] + sides[1]
}

// getRibbonForBow returns the amount of ribbon needed for a bow
func getRibbonForBow(l, w, h int) int {
	return l * w * h
}
