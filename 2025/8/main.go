package main

import (
	"fmt"
	"slices"
	"sort"
	"time"

	"utils"
)

const day = 8

type junctionBoxes struct {
	p1, p2, distance int
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
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	var x1, x2, y1, y2, z1, z2 int
	var jBoxes []junctionBoxes

	for i := 0; i < len(lines)-1; i++ {
		_, _ = fmt.Sscanf(lines[i], "%d,%d,%d", &x1, &y1, &z1)
		for j := i + 1; j < len(lines); j++ {
			_, _ = fmt.Sscanf(lines[j], "%d,%d,%d", &x2, &y2, &z2)
			distance := ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)) + ((z2 - z1) * (z2 - z1))
			jBoxes = append(jBoxes, junctionBoxes{i, j, distance})
		}
	}

	sort.Slice(jBoxes, func(i, j int) bool {
		return jBoxes[i].distance < jBoxes[j].distance
	})

	if len(jBoxes) >= 1000 {
		jBoxes = jBoxes[:1000]
	} else {
		jBoxes = jBoxes[:10]
	}

	circuits := make([][]int, 0)
	var foundP1, foundP2 int
	for i := 0; i < len(jBoxes); i++ {
		foundP1, foundP2 = -1, -1
		for j := 0; j < len(circuits); j++ {
			for _, b := range circuits[j] {
				if b == jBoxes[i].p1 {
					foundP1 = j
					break
				}
			}
			for _, b := range circuits[j] {
				if b == jBoxes[i].p2 {
					foundP2 = j
					break
				}
			}
		}

		switch {
		case foundP1 == -1 && foundP2 == -1:
			circuits = append(circuits, []int{jBoxes[i].p1, jBoxes[i].p2})
		case foundP1 == foundP2:
			continue
		case foundP1 != -1 && foundP2 == -1:
			circuits[foundP1] = append(circuits[foundP1], jBoxes[i].p2)
		case foundP1 == -1 && foundP2 != -1:
			circuits[foundP2] = append(circuits[foundP2], jBoxes[i].p1)
		case foundP1 != -1 && foundP2 != -1 && foundP1 != foundP2:
			circuits[foundP1] = append(circuits[foundP1], circuits[foundP2]...) // add p2 on p1
			circuits = append(circuits[:foundP2], circuits[foundP2+1:]...)      // remove p2
		}

	}
	slices.SortFunc(circuits, func(c1, c2 []int) int {
		return len(c2) - len(c1)
	})
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func solutionB() int {
	solution := 0
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	var x1, x2, y1, y2, z1, z2 int
	var jBoxes []junctionBoxes

	for i := 0; i < len(lines)-1; i++ {
		_, _ = fmt.Sscanf(lines[i], "%d,%d,%d", &x1, &y1, &z1)
		for j := i + 1; j < len(lines); j++ {
			_, _ = fmt.Sscanf(lines[j], "%d,%d,%d", &x2, &y2, &z2)
			distance := ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)) + ((z2 - z1) * (z2 - z1))
			jBoxes = append(jBoxes, junctionBoxes{i, j, distance})
		}
	}

	sort.Slice(jBoxes, func(i, j int) bool {
		return jBoxes[i].distance < jBoxes[j].distance
	})

	circuits := make([][]int, 0)
	var foundP1, foundP2 int
	for i := 0; i < len(jBoxes); i++ {
		foundP1, foundP2 = -1, -1
		for j := 0; j < len(circuits); j++ {
			for _, b := range circuits[j] {
				if b == jBoxes[i].p1 {
					foundP1 = j
					break
				}
			}
			for _, b := range circuits[j] {
				if b == jBoxes[i].p2 {
					foundP2 = j
					break
				}
			}
		}

		switch {
		case foundP1 == -1 && foundP2 == -1:
			circuits = append(circuits, []int{jBoxes[i].p1, jBoxes[i].p2})
		case foundP1 == foundP2:
			continue
		case foundP1 != -1 && foundP2 == -1:
			circuits[foundP1] = append(circuits[foundP1], jBoxes[i].p2)
		case foundP1 == -1 && foundP2 != -1:
			circuits[foundP2] = append(circuits[foundP2], jBoxes[i].p1)
		case foundP1 != -1 && foundP2 != -1 && foundP1 != foundP2:
			circuits[foundP1] = append(circuits[foundP1], circuits[foundP2]...) // add p2 on p1
			circuits = append(circuits[:foundP2], circuits[foundP2+1:]...)      // remove p2
		}

		_, _ = fmt.Sscanf(lines[jBoxes[i].p1], "%d,%d,%d", &x1, &y1, &z1)
		_, _ = fmt.Sscanf(lines[jBoxes[i].p2], "%d,%d,%d", &x2, &y2, &z2)
		solution = x1 * x2

	}

	return solution
}
