package main

import (
	"fmt"
	"math"
	"slices"
	"time"
	"utils"
)

const day = 9

var distances = map[string]int{}

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

	cities := parseLine(lines)
	solution = findShortestRoute(cities)

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	cities := parseLine(lines)
	solution = findLongestRoute(cities)

	return solution
}

// parseLine parses the input lines and populates the distances map
func parseLine(lines []string) []string {
	cities := make([]string, 0)
	var to, from string
	var distance int

	for _, line := range lines {
		_, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if !slices.Contains(cities, from) {
			cities = append(cities, from)
		}

		if !slices.Contains(cities, to) {
			cities = append(cities, to)
		}

		key := makeKey(from, to)
		distances[key] = distance
	}
	return cities
}

// makeKey creates a consistent key for the distance map
func makeKey(city1, city2 string) string {
	if city1 < city2 {
		return string(city1) + ":" + string(city2)
	}
	return string(city2) + ":" + string(city1)
}

// findShortestRoute finds the shortest route between all cities
func findShortestRoute(cities []string) int {
	permutations := permute(cities)
	shortest := math.MaxInt32

	for _, route := range permutations {
		distance := calculateRouteDistance(route)
		if distance < shortest {
			shortest = distance
		}
	}

	return shortest
}

// findLongestRoute finds the longest route between all cities
func findLongestRoute(cities []string) int {
	permutations := permute(cities)
	longest := math.MinInt32

	for _, route := range permutations {
		distance := calculateRouteDistance(route)
		if distance > longest {
			longest = distance
		}
	}

	return longest
}

// permute generates all permutations of a slice of strings
func permute(cities []string) [][]string {
	var helper func([]string, int)
	var res [][]string

	helper = func(cities []string, n int) {
		if n == 1 {
			tmp := make([]string, len(cities))
			copy(tmp, cities)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(cities, n-1)
				if n%2 == 1 {
					cities[0], cities[n-1] = cities[n-1], cities[0]
				} else {
					cities[i], cities[n-1] = cities[n-1], cities[i]
				}
			}
		}
	}

	helper(cities, len(cities))
	return res
}

// calculateRouteDistance calculates the distance of a route
func calculateRouteDistance(route []string) int {
	distance := 0
	for i := 0; i < len(route)-1; i++ {
		key := makeKey(route[i], route[i+1])
		distance += distances[key]
	}
	return distance
}
