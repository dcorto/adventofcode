package main

import (
	"fmt"
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
			lol := Replacement{from: r[0], to: r[1]}
			replacements = append(replacements, lol)
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
	var solution = 0
	return solution
}
