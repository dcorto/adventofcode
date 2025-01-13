package main

import (
	"fmt"
	"time"
	"utils"
)

const day = 8

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

	chars := 0
	memory := 0
	for _, line := range lines {
		chars += len(line)
		memory += parseString(line)
	}

	solution = chars - memory

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
		solution += len(encodeString(line)) - len(line)
	}

	return solution
}

// parseString returns the number of characters in the string literal (excluding the surrounding double quotes)
func parseString(s string) int {
	count := 0
	i := 1
	for i < len(s)-1 {
		if s[i] == '\\' { // backslash
			if i+1 >= len(s)-1 {
				break // safety check
			}
			switch s[i+1] {
			case '\\', '"':
				count++
				i += 2 // skip both the backslash and the escaped char
			case 'x':
				if i+3 >= len(s)-1 {
					break // safety check for hex sequence
				}
				count++
				i += 4 // skip \x and the two hex digits
			}
		} else {
			count++
			i++
		}
	}

	return count
}

// encodeString returns the encoded version of the string
func encodeString(s string) string {
	newString := "\"" // start with double quote
	for _, c := range s {
		if c == '\\' {
			newString += "\\\\"
		} else if c == '"' {
			newString += "\\\""
		} else {
			newString += string(c)
		}

	}

	newString += "\""
	return newString
}
