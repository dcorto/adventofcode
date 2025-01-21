package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 14

const RaceDuration = 2503

type Reindeer struct {
	Name            string // name of the reindeer
	Speed           int    // the speed at which it can fly
	FlightTime      int    // how long it can fly at its speed before rest
	RestTime        int    // how long it must rest before it can fly again
	RestRemaining   int    // time remaining to rest
	FlightRemaining int    // time remaining to fly
	Distance        int    // distance traveled
	Points          int    // points earned
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
	var deer []Reindeer
	for _, line := range lines {
		reindeer := parseLine(line)
		deer = append(deer, reindeer)
	}

	doRace(deer)

	furthestDistance := 0
	for _, d := range deer {
		furthestDistance = max(furthestDistance, d.Distance)
	}
	solution = furthestDistance

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}
	var deer []Reindeer
	for _, line := range lines {
		reindeer := parseLine(line)
		deer = append(deer, reindeer)
	}

	doRaceWithPoints(deer)

	maxPoints := 0
	for _, d := range deer {
		maxPoints = max(maxPoints, d.Points)
	}
	solution = maxPoints

	return solution
}

// parseLine parses a line of input and returns the person, neighbor and happiness units
func parseLine(line string) (reindeer Reindeer) {
	var name string
	var speed, duration, rest int
	line = strings.TrimSuffix(line, ".") // remove trailing dot
	_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds", &name, &speed, &duration, &rest)
	if err != nil {
		fmt.Println("Error:", err)
		return Reindeer{}
	}
	return Reindeer{
		Name:            name,
		Speed:           speed,
		FlightTime:      duration,
		FlightRemaining: duration,
		RestTime:        rest,
		RestRemaining:   rest,
	}
}

// doRace simulates the race of all reindeer
func doRace(deer []Reindeer) {
	for i := 0; i < RaceDuration; i++ {
		for k, d := range deer {
			if deer[k].FlightRemaining > 0 {
				deer[k].FlightRemaining--
				if deer[k].FlightRemaining == 0 {
					deer[k].RestRemaining = deer[k].RestTime
				}

				deer[k].Distance += d.Speed
			} else {
				deer[k].RestRemaining--
				if deer[k].RestRemaining == 0 {
					deer[k].FlightRemaining = deer[k].FlightTime
				}
			}
		}
	}
}

// doRaceWithPoints simulates the race of all reindeer with points system
func doRaceWithPoints(deer []Reindeer) {
	for i := 0; i < RaceDuration; i++ {
		for k, d := range deer {
			if deer[k].FlightRemaining > 0 {
				deer[k].FlightRemaining--
				if deer[k].FlightRemaining == 0 {
					deer[k].RestRemaining = deer[k].RestTime
				}
				deer[k].Distance += d.Speed
			} else {
				deer[k].RestRemaining--
				if deer[k].RestRemaining == 0 {
					deer[k].FlightRemaining = deer[k].FlightTime
				}
			}
		}
		awardPoints(deer)
	}
}

// awardPoints awards points to the leading reindeer
func awardPoints(deer []Reindeer) {
	maxDist := 0

	for _, d := range deer {
		if d.Distance > maxDist {
			maxDist = d.Distance
		}
	}

	for i := range deer {
		if deer[i].Distance == maxDist {
			deer[i].Points++
		}
	}
}
