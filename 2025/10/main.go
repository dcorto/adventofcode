package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"utils"
)

const day = 10

type Counter []int
type Machine struct {
	goal    Counter
	buttons [][]int
	counter Counter
}

type ButtonCombination struct {
	counter        Counter
	pressedButtons int
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
		return 0
	}

	machines := generateMachines(lines)
	for _, m := range machines {
		var combinations = allCombinations(m.buttons, len(m.counter))
		var n = m.goal.solveA(combinations)
		solution += n
	}

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	machines := generateMachines(lines)
	for _, m := range machines {
		var combinations = allCombinations(m.buttons, len(m.counter))
		var n, _ = m.counter.solveB(combinations)
		solution += n
	}

	return solution
}

func generateMachines(lines []string) []Machine {
	var machines []Machine
	for _, line := range lines {
		var fields = strings.Fields(line)
		var goal []int
		for i := 1; i < len(fields[0])-1; i++ {
			if fields[0][i] == '#' {
				goal = append(goal, 1)
			} else {
				goal = append(goal, 0)
			}
		}
		var buttons [][]int
		for j := 1; j < len(fields)-1; j++ {
			buttons = append(buttons, buildIntList(fields[j][1:len(fields[j])-1]))
		}
		var lastField = fields[len(fields)-1]
		var counter = buildIntList(lastField[1 : len(lastField)-1])
		machines = append(machines, Machine{goal, buttons, counter})
	}
	return machines
}

func buildIntList(s string) []int {
	var res []int
	var l = strings.Split(s, ",")
	for _, v := range l {
		var n, _ = strconv.Atoi(v)
		res = append(res, n)
	}
	return res
}

func allCombinations(buttons [][]int, m int) []ButtonCombination {
	var nbButtons = len(buttons)
	if nbButtons == 0 {
		return []ButtonCombination{{counter: make([]int, m), pressedButtons: 0}}
	}

	var res = make([]ButtonCombination, 0, 1<<nbButtons)
	for n := 0; n < (1 << nbButtons); n++ {
		var counter = make([]int, m)
		var pressedButtons = 0
		for j := 0; j < nbButtons; j++ {
			if (n & (1 << j)) != 0 {
				pressedButtons++
				for _, idx := range buttons[j] {
					counter[idx]++
				}
			}
		}
		res = append(res, ButtonCombination{counter, pressedButtons})
	}
	return res
}

func newCounter(size int) Counter {
	return make([]int, size)
}

func (c Counter) solveA(combinations []ButtonCombination) int {
	var res = math.MaxInt
	for _, comb := range combinations {
		if comb.counter.isModulo2(c) {
			if comb.pressedButtons < res {
				res = comb.pressedButtons
			}
		}
	}
	return res
}

func (c Counter) solveB(combinations []ButtonCombination) (int, bool) {
	if c.isZero() {
		return 0, true
	}

	var res = math.MaxInt
	for _, comb := range combinations {
		if !comb.counter.smallerOrEqual(c) {
			continue
		}
		if !comb.counter.isModulo2(c) {
			continue
		}

		var nextCounter = newCounter(len(c))
		for i := 0; i < len(c); i++ {
			nextCounter[i] = (c[i] - comb.counter[i]) / 2
		}
		rec, ok := nextCounter.solveB(combinations)
		if !ok {
			continue
		}

		if n := 2*rec + comb.pressedButtons; n < res {
			res = n
		}
	}

	if res < math.MaxInt {
		return res, true
	}

	return 0, false
}

func (c Counter) isModulo2(b Counter) bool {
	for i := 0; i < len(c); i++ {
		if c[i]%2 != b[i]%2 {
			return false
		}
	}
	return true
}

func (c Counter) isZero() bool {
	for i := 0; i < len(c); i++ {
		if c[i] != 0 {
			return false
		}
	}
	return true
}

func (c Counter) smallerOrEqual(b Counter) bool {
	for i := 0; i < len(c); i++ {
		if c[i] > b[i] {
			return false
		}
	}
	return true
}
