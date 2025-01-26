package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"utils"
)

const day = 16

type Aunt struct {
	ID          int
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

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

	desiredAunt := Aunt{
		ID:          -1,
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	aunts := parseLines(lines)

	similarity := make(map[int]int)
	for _, a := range aunts {
		similarity[a.ID] = mesureSimilarity(a, desiredAunt)
	}

	maxSimilarity := 0
	for id, s := range similarity {
		if s > maxSimilarity {
			maxSimilarity = s
			desiredAunt.ID = id
		}
	}

	solution = desiredAunt.ID

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	desiredAunt := Aunt{
		ID:          -1,
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	aunts := parseLines(lines)

	auntsMap := make(map[int]Aunt)
	for _, a := range aunts {
		auntsMap[a.ID] = a
	}

	similarity := make(map[int]int)
	for _, a := range auntsMap {
		similarity[a.ID] = mesureSimilarityRange(a, desiredAunt)
	}

	maxSimilarity := 0
	for id, s := range similarity {
		if s > maxSimilarity {
			maxSimilarity = s
			desiredAunt.ID = id
		}
	}

	solution = desiredAunt.ID

	return solution
}

// parseLines returns a slice of Aunts from a slice of strings
func parseLines(lines []string) []Aunt {
	var aunts []Aunt

	for _, line := range lines {

		a := Aunt{
			ID:          0, // always at least 1
			children:    -1,
			cats:        -1,
			samoyeds:    -1,
			pomeranians: -1,
			akitas:      -1,
			vizslas:     -1,
			goldfish:    -1,
			trees:       -1,
			cars:        -1,
			perfumes:    -1,
		}

		a.ID, _ = strconv.Atoi(strings.Split(strings.Split(line, ": ")[0], " ")[1])

		// trim the "Sue n: " from the beginning of each line
		line = strings.TrimPrefix(line, "Sue "+strconv.Itoa(a.ID)+": ")
		parts := strings.Split(line, ", ")

		for _, part := range parts[0:] {
			switch strings.Split(part, ": ")[0] {
			case "children":
				a.children, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "cats":
				a.cats, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "samoyeds":
				a.samoyeds, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "pomeranians":
				a.pomeranians, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "akitas":
				a.akitas, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "vizslas":
				a.vizslas, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "goldfish":
				a.goldfish, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "trees":
				a.trees, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "cars":
				a.cars, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			case "perfumes":
				a.perfumes, _ = strconv.Atoi(strings.Split(part, ": ")[1])
				continue
			default:
				continue
			}
		}

		aunts = append(aunts, a)
	}

	return aunts
}

// mesureSimilarity returns the number of similar properties between two Aunts
func mesureSimilarity(a, b Aunt) int {
	similarity := 0

	if a.children == b.children {
		similarity++
	}
	if a.cats == b.cats {
		similarity++
	}
	if a.samoyeds == b.samoyeds {
		similarity++
	}
	if a.pomeranians == b.pomeranians {
		similarity++
	}
	if a.akitas == b.akitas {
		similarity++
	}
	if a.vizslas == b.vizslas {
		similarity++
	}
	if a.goldfish == b.goldfish {
		similarity++
	}
	if a.trees == b.trees {
		similarity++
	}
	if a.cars == b.cars {
		similarity++
	}
	if a.perfumes == b.perfumes {
		similarity++
	}

	return similarity
}

// mesureSimilarityRange returns the number of similar properties between two Aunts
func mesureSimilarityRange(a, b Aunt) int {
	similarity := 0

	if a.children == 3 {
		similarity++
	}
	if a.cats > 7 {
		similarity++
	}
	if a.samoyeds == 2 {
		similarity++
	}
	if a.pomeranians < 3 {
		similarity++
	}
	if a.akitas == 0 {
		similarity++
	}
	if a.vizslas == 0 {
		similarity++
	}
	if a.goldfish < 3 {
		similarity++
	}
	if a.trees > 3 {
		similarity++
	}
	if a.cars == 2 {
		similarity++
	}
	if a.perfumes == 1 {
		similarity++
	}

	return similarity
}
