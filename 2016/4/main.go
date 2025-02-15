package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"utils"
)

const day = 4

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

	var finalSum int
	for _, line := range lines {
		re := regexp.MustCompile(`([a-z-]+)-(\d+)\[([a-z]+)]`)
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		name := matches[1]
		sectorID, _ := strconv.Atoi(matches[2])
		entryChecksum := matches[3]

		charCount := make(map[rune]int)
		for _, char := range name {
			if char != '-' {
				charCount[char]++
			}
		}

		type kv struct {
			Key   rune
			Value int
		}

		var sortedChars []kv
		for k, v := range charCount {
			sortedChars = append(sortedChars, kv{k, v})
		}

		sort.Slice(sortedChars, func(i, j int) bool {
			if sortedChars[i].Value == sortedChars[j].Value {
				return sortedChars[i].Key < sortedChars[j].Key
			}
			return sortedChars[i].Value > sortedChars[j].Value
		})

		var checksumBuilder strings.Builder
		for i := 0; i < 5; i++ {
			checksumBuilder.WriteRune(sortedChars[i].Key)
		}
		checksum := checksumBuilder.String()

		if checksum == entryChecksum {
			finalSum += sectorID
		}
	}

	solution = finalSum
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
		re := regexp.MustCompile(`([a-z-]+)-(\d+)\[([a-z]+)]`)
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		name := matches[1]
		sectorID, _ := strconv.Atoi(matches[2])

		decryptedName := ""
		for _, char := range name {
			decryptedName += string(shift(char, sectorID))
		}

		if strings.Contains(decryptedName, "north") && strings.Contains(decryptedName, "pole") {
			solution = sectorID
		}
	}

	return solution
}

// shift shifts a character by count
func shift(c rune, count int) rune {
	if c == '-' {
		return ' '
	}

	count = count % 26
	ac := int(c) - 'a'
	ac += count
	ac = ac % 26
	return rune(ac + 'a')
}
