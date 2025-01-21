package main

import (
	"fmt"
	"math"
	"strings"
	"time"
	"utils"
)

const day = 13

// Person is a type for a person
type Person string

// Happiness is a type for a map of Person and their happiness
type Happiness map[Person]map[Person]int

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

	happiness := make(Happiness)
	people := map[Person]struct{}{}

	for _, line := range lines {
		person, neighbor, units := parseLine(line)

		// add to happiness map
		people[person] = struct{}{}
		people[neighbor] = struct{}{}

		// init happiness map
		if happiness[person] == nil {
			happiness[person] = make(map[Person]int)
		}
		happiness[person][neighbor] = units
	}

	// Convert map to slice for permutations
	var peopleSlice []Person
	for person := range people {
		peopleSlice = append(peopleSlice, person)
	}

	permutationsOfPeople := permute(peopleSlice)
	solution = calcMaxHappiness(happiness, permutationsOfPeople)

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	happiness := make(Happiness)
	people := map[Person]struct{}{}

	for _, line := range lines {
		person, neighbor, units := parseLine(line)

		// add to happiness map
		people[person] = struct{}{}
		people[neighbor] = struct{}{}

		// init happiness map
		if happiness[person] == nil {
			happiness[person] = make(map[Person]int)
		}
		happiness[person][neighbor] = units
	}

	// Convert map to slice for permutations
	var peopleSlice []Person
	for person := range people {
		peopleSlice = append(peopleSlice, person)
	}

	happiness, peopleSlice = addMe(happiness, peopleSlice)

	permutationsOfPeople := permute(peopleSlice)
	solution = calcMaxHappiness(happiness, permutationsOfPeople)

	return solution
}

// parseLine parses a line of input and returns the person, neighbor and happiness units
func parseLine(line string) (Person, Person, int) {
	var person, neighbor, direction string
	var units int

	line = strings.TrimSuffix(line, ".") // remove trailing dot
	_, err := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &person, &direction, &units, &neighbor)
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", 0
	}

	if direction == "lose" {
		units = -units
	}

	return Person(person), Person(neighbor), units
}

// permute returns all possible permutations of people
func permute(people []Person) [][]Person {
	var helper func([]Person, int)
	var res [][]Person

	helper = func(people []Person, n int) {
		if n == 1 {
			tmp := make([]Person, len(people))
			copy(tmp, people)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(people, n-1)
				if n%2 == 1 {
					people[0], people[n-1] = people[n-1], people[0]
				} else {
					people[i], people[n-1] = people[n-1], people[i]
				}
			}
		}
	}

	helper(people, len(people))
	return res
}

// calcMaxHappiness calculates the maximum happiness for a given set of people
func calcMaxHappiness(h Happiness, people [][]Person) int {
	maxHappiness := math.MinInt
	for _, p := range people {
		total := 0

		for i := 0; i < len(p); i++ {
			curr := p[i]
			l := p[(i-1+len(p))%len(p)]
			r := p[(i+1)%len(p)]

			// add happiness from both neighbors
			total += h[curr][l]
			total += h[curr][r]
		}

		if total > maxHappiness {
			maxHappiness = total
		}
	}

	return maxHappiness
}

// addMe adds a new person to the happiness map and people slice
func addMe(h Happiness, ps []Person) (Happiness, []Person) {
	me := Person("me")
	// Add me to the slice
	nps := append(ps, me)

	// Initialize my happiness map
	h[me] = make(map[Person]int)

	// Add 0 happiness for all relationships involving me
	for _, p := range nps {
		if p != me {
			// me -> others
			h[me][p] = 0

			// others -> me
			if h[p] == nil {
				h[p] = make(map[Person]int)
			}
			h[p][me] = 0
		}
	}

	return h, nps
}
