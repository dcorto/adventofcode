package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 3

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

	directions := strings.Split(input, "")
	visited := make(map[[2]int]bool)
	position := [2]int{0, 0}
	housesVisited := 1
	visited[position] = true //initial position is visited

	for _, d := range directions {
		position = move(position, d)
		if !checkIfPositionIsVisited(visited, position) {
			visited[position] = true
			housesVisited++
		}
	}

	solution = housesVisited

	return solution
}

func solutionB() int {
	var solution = 0

	input, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	directions := strings.Split(input, "")
	visited := make(map[[2]int]bool)
	positionSanta := [2]int{0, 0}
	positionRoboSanta := [2]int{0, 0}
	housesVisited := 1

	visited[positionSanta] = true //initial position is visited (by Santa and RoboSanta)

	for i, d := range directions {
		if i%2 == 0 { //Santa moves
			positionSanta = move(positionSanta, d)
			if !checkIfPositionIsVisited(visited, positionSanta) {
				visited[positionSanta] = true
				housesVisited++
			}
			continue
		}
		// RoboSanta moves
		positionRoboSanta = move(positionRoboSanta, d)
		if !checkIfPositionIsVisited(visited, positionRoboSanta) {
			visited[positionRoboSanta] = true
			housesVisited++
		}
	}

	solution = housesVisited

	return solution
}

// checkIfPositionIsVisited by Santa or RoboSanta
func checkIfPositionIsVisited(visited map[[2]int]bool, position [2]int) bool {
	if _, ok := visited[position]; ok {
		return true
	}
	return false
}

// move Santa or RoboSanta to a new position
func move(position [2]int, d string) [2]int {
	//move north
	if d == "^" {
		return [2]int{position[0] + 1, position[1]}
	}

	//move east
	if d == ">" {
		return [2]int{position[0], position[1] + 1}
	}

	//move south
	if d == "v" {
		return [2]int{position[0] - 1, position[1]}
	}

	//move west
	if d == "<" {
		return [2]int{position[0], position[1] - 1}
	}

	return [2]int{position[0], position[1]}
}
