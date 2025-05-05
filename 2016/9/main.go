package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"utils"
)

const day = 9

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

	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	decompressedLength := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			markerEnd := strings.IndexByte(data[i:], ')')
			marker := data[i+1 : i+markerEnd]
			parts := strings.Split(marker, "x")
			chars, _ := strconv.Atoi(parts[0])
			repeat, _ := strconv.Atoi(parts[1])

			decompressedLength += chars * repeat
			i += markerEnd + chars
		} else {
			decompressedLength++
		}
	}

	solution = decompressedLength

	return solution
}

func solutionB() int {
	var solution = 0
	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}
	dataBytes := []byte(data) // Work with byte slices
	solution = calculateDecompressedLengthV2(dataBytes)

	return solution
}

// calculateDecompressedLengthV2 calculates length recursively for Part B
func calculateDecompressedLengthV2(data []byte) int {
	length := 0
	i := 0
	for i < len(data) {
		if data[i] == '(' {
			markerEndInSubstring := bytes.IndexByte(data[i:], ')')
			// Assuming valid input, markerEndInSubstring will not be -1
			markerEnd := i + markerEndInSubstring

			marker := string(data[i+1 : markerEnd])
			parts := strings.Split(marker, "x")
			// Assuming valid input, len(parts) == 2 and Atoi succeeds

			chars, _ := strconv.Atoi(parts[0])
			repeat, _ := strconv.Atoi(parts[1])

			// Define the segment that needs recursive decompression
			segmentStart := markerEnd + 1
			segmentEnd := segmentStart + chars
			// Boundary check (important!)
			if segmentEnd > len(data) {
				segmentEnd = len(data)
			}
			segment := data[segmentStart:segmentEnd]

			// Recursively calculate the length of the segment
			recursiveLength := calculateDecompressedLengthV2(segment)

			// Add the repeated length to the total
			length += recursiveLength * repeat

			// Advance the index past the marker and the segment it referred to
			i = segmentEnd

		} else {
			// Regular character, just count it
			length++
			i++
		}
	}
	return length
}
