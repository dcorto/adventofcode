package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"utils"
)

const day = 7

var Operations = [6]string{"ASSIGN", "AND", "OR", "LSHIFT", "RSHIFT", "NOT"}

type Instruction struct {
	Operation string
	Inputs    []string
	Output    string
}

type Wire struct {
	Value    int
	Assigned bool
}

var instructions []Instruction
var wires map[string]*Wire

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

	generateInstructions(lines)

	// Run & execute instructions
	wires = make(map[string]*Wire)
	runInstructions()

	if wire, exists := wires["a"]; exists {
		solution = wire.Value
	}

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	generateInstructions(lines)

	// Keep a copy of original instructions
	originalInstructions := make([]Instruction, len(instructions))
	copy(originalInstructions, instructions)

	// Execute instructions
	wires = make(map[string]*Wire)
	runInstructions()

	overrideB := wires["a"].Value

	// Reset for second run
	wires = make(map[string]*Wire)
	wires["b"] = &Wire{Value: overrideB, Assigned: true}

	// Reset instructions, excluding the one setting 'b'
	instructions = []Instruction{}
	for _, inst := range originalInstructions {
		if inst.Output != "b" {
			instructions = append(instructions, inst)
		}
	}

	runInstructions()

	if wire, exists := wires["a"]; exists {
		solution = wire.Value
	}

	return solution
}

// runInstructions runs the instructions left to be executed
func runInstructions() {
	for len(instructions) > 0 {
		var remainingInstructions []Instruction

		for _, inst := range instructions {
			if inst.Output == "b" && wires["b"] != nil && wires["b"].Assigned {
				// Skip instructions trying to assign to 'b' if it's already set (for part B)
				continue
			}

			if !executeInstruction(inst) {
				// Instruction couldn't be executed yet
				remainingInstructions = append(remainingInstructions, inst)
			}
		}

		if len(remainingInstructions) == len(instructions) {
			panic("No progress made, might be a circular dependency")
		}

		instructions = remainingInstructions
	}
}

// executeInstruction executes the given instruction and returns true if successful
func executeInstruction(inst Instruction) bool {
	// given an instruction, execute it and return true if successful
	switch inst.Operation {
	case "ASSIGN":
		if val, ok := getValue(inst.Inputs[0]); ok {
			wires[inst.Output] = &Wire{Value: val, Assigned: true}
			return true
		}
	case "NOT":
		if val, ok := getValue(inst.Inputs[0]); ok {
			wires[inst.Output] = &Wire{Value: ^val & 0xFFFF, Assigned: true}
			return true
		}
	default: // AND, OR, LSHIFT, RSHIFT
		if len(inst.Inputs) != 2 {
			panic("Invalid number of inputs for operation " + inst.Operation)
		}
		val1, ok1 := getValue(inst.Inputs[0])
		val2, ok2 := getValue(inst.Inputs[1])
		if ok1 && ok2 {
			result := doOperation(inst.Operation, val1, val2)
			wires[inst.Output] = &Wire{Value: result, Assigned: true}
			return true
		}
	}
	return false
}

// generateInstructions parses the input lines and generates the instructions
func generateInstructions(lines []string) {
	for _, line := range lines {
		instruction := Instruction{}

		// Parse the line
		parts := strings.Split(line, " -> ")

		// Gate
		isGate := false
		op := ""
		for i, operation := range Operations {
			if strings.Contains(parts[0], operation) {
				op = Operations[i]
				isGate = true
				break
			}
		}

		if isGate {
			output := parts[1]
			a, b := "", ""
			if op == "NOT" {
				a = strings.TrimSpace(strings.ReplaceAll(parts[0], op, ""))
			} else {
				a, b = strings.TrimSpace(strings.Split(parts[0], op)[0]), strings.TrimSpace(strings.Split(parts[0], op)[1])
			}

			instruction.Inputs = []string{a, b}
			instruction.Operation = op
			instruction.Output = output
			instructions = append(instructions, instruction)
			continue
		}

		// Wire
		value, output := parts[0], parts[1]
		instruction.Inputs = []string{value}
		instruction.Operation = Operations[0] // ASSIGN
		instruction.Output = output

		instructions = append(instructions, instruction)
	}
}

// getValue gets the value of a given input string. could be an existing wire or a number.
// Returns false if the input string is not a wire or number (error)
func getValue(input string) (int, bool) {
	if val, err := strconv.ParseUint(input, 10, 16); err == nil {
		return int(val), true
	}

	if wire, exists := wires[input]; exists && wire.Assigned {
		return wire.Value, true
	}
	return 0, false
}

// doOperation does the operation on the two inputs
func doOperation(op string, a, b int) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "LSHIFT":
		return a << b
	case "RSHIFT":
		return a >> b
	case "NOT":
		return ^a & 0xFFFF // Ensures the result is 16-bit
	default:
		panic("Invalid operation " + op)
	}
}
