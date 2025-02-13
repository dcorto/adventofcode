package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"utils"
)

const day = 19

type Replacement struct {
	from string
	to   string
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

	isMolecule := false
	var molecule = ""
	var replacements []Replacement
	for _, line := range lines {
		if line == "" {
			isMolecule = true
			continue
		}

		if isMolecule {
			molecule = line
		} else {
			r := strings.Split(line, " => ")
			replacement := Replacement{from: r[0], to: r[1]}
			replacements = append(replacements, replacement)
		}
	}

	distinctMolecules := make(map[string]struct{})
	for _, r := range replacements {
		from := r.from
		to := r.to
		for i := 0; i < len(molecule)-len(from)+1; i++ {
			if molecule[i:i+len(from)] == from {
				newMolecule := molecule[:i] + to + molecule[i+len(from):]
				distinctMolecules[newMolecule] = struct{}{}
			}
		}
	}

	solution = len(distinctMolecules)

	return solution
}

func solutionB() int {
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	isMolecule := false
	var molecule string
	var replacements = map[string][]string{}
	for _, line := range lines {
		if line == "" {
			isMolecule = true
			continue
		}

		if isMolecule {
			molecule = line
		} else {
			r := strings.Split(line, " => ")
			from, to := r[0], r[1]
			replacements[from] = append(replacements[from], to)
		}
	}

	randomReplacements := getRandomMappings(replacements)
	steps := 0

	for molecule != "e" {
		replaced := false
		for k, v := range randomReplacements {
			if strings.Contains(molecule, k) {
				molecule = strings.Replace(molecule, k, v, 1)
				steps++
				replaced = true
				break
			}
		}
		if !replaced {
			return steps
		}
	}

	return steps
}

// getRandomMappings returns a map with the keys of the input map in a random order
func getRandomMappings(input map[string][]string) map[string]string {
	reverse := make(map[string]string)
	for start, results := range input {
		for _, res := range results {
			reverse[res] = start
		}
	}
	keys := make([]string, 0, len(reverse))
	for k := range reverse {
		keys = append(keys, k)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	sortedRandom := make(map[string]string)
	for _, k := range keys {
		sortedRandom[k] = reverse[k]
	}
	return sortedRandom
}
