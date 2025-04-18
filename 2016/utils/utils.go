package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadLinesFromFile reads a file and returns a slice of strings one for each line
func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return lines, nil
}

// ReadFromFile reads a file and returns it's content as a string
func ReadFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	content := string(data)
	return content, nil
}

// Atoi cast a string to an integer
func Atoi(line string) int {
	n, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Error casting string to int:", err)
		return 0
	}
	return n
}
