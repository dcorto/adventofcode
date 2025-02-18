package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"utils"
)

const day = 7

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

	var tlsCount = 0
	for _, ip := range lines {
		if supportsTLS(ip) {
			tlsCount++
		}
	}

	solution = tlsCount
	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var sslCount = 0
	for _, ip := range lines {
		if supportsSSL(ip) {
			sslCount++
		}
	}

	solution = sslCount

	return solution
}

// supportsTLS checks if an IP supports TLS
func supportsTLS(ip string) bool {
	bracketPattern := regexp.MustCompile(`\[(.*?)]`)
	hypernetSequences := bracketPattern.FindAllString(ip, -1)
	for _, seq := range hypernetSequences {
		if containsABBA(seq[1 : len(seq)-1]) {
			return false
		}
	}
	nonHypernetSequences := bracketPattern.Split(ip, -1)
	for _, seq := range nonHypernetSequences {
		if containsABBA(seq) {
			return true
		}
	}
	return false
}

// supportsSSL checks if an IP supports SSL
func supportsSSL(ip string) bool {
	bracketPattern := regexp.MustCompile(`\[(.*?)]`)
	hypernetSequences := bracketPattern.FindAllString(ip, -1)
	nonHypernetSequences := bracketPattern.Split(ip, -1)

	var abas []string
	for _, seq := range nonHypernetSequences {
		abas = append(abas, findABAs(seq)...)
	}

	for _, seq := range hypernetSequences {
		for _, aba := range abas {
			bab := string([]byte{aba[1], aba[0], aba[1]})
			if strings.Contains(seq, bab) {
				return true
			}
		}
	}
	return false
}

// containsABBA checks if a string contains an ABBA pattern
func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

// findABAs finds all ABAs in a string
func findABAs(s string) []string {
	var abas []string
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			abas = append(abas, s[i:i+3])
		}
	}
	return abas
}
