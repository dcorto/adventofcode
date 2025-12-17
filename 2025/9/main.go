package main

import (
	"fmt"
	"math"
	"time"

	"utils"
)

const day = 9

type Point struct {
	x, y int
}

type Segment struct {
	p1, p2 Point
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
	points := getPoints(lines)

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			area := int((math.Abs(float64(p2.x-p1.x)) + 1) * (math.Abs(float64(p2.y-p1.y)) + 1))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func solutionB() int {
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	redTiles := getPoints(lines)
	if len(redTiles) < 2 {
		return 0
	}

	// Build segments
	var segments []Segment
	for i := 0; i < len(redTiles); i++ {
		segments = append(segments, Segment{redTiles[i], redTiles[(i+1)%len(redTiles)]})
	}

	maxArea := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			p1 := redTiles[i]
			p2 := redTiles[j]

			// Calculate area
			width := int(math.Abs(float64(p2.x-p1.x))) + 1
			height := int(math.Abs(float64(p2.y-p1.y))) + 1
			area := width * height

			if area > maxArea {
				// Check if valid
				if isValid(p1, p2, segments, redTiles) {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// getPoints returns points from line
func getPoints(lines []string) []Point {
	var points []Point
	for _, line := range lines {
		var x, y int
		_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, Point{x, y})
	}
	return points
}

// isValid checks if Point is valid
func isValid(p1, p2 Point, segments []Segment, polygon []Point) bool {
	for _, seg := range segments {
		if intersect(seg, p1, p2) {
			return false
		}
	}

	minX := float64(min(p1.x, p2.x))
	maxX := float64(max(p1.x, p2.x))
	minY := float64(min(p1.y, p2.y))
	maxY := float64(max(p1.y, p2.y))

	centerX := (minX + maxX) / 2.0
	centerY := (minY + maxY) / 2.0

	return isInside(centerX, centerY, polygon)
}

// intersect checks if a segment intersects the strict interior of the rectangle defined by p1, p2
func intersect(segment Segment, p1, p2 Point) bool {
	// Rectangle interior bounds
	minX := min(p1.x, p2.x) + 1
	maxX := max(p1.x, p2.x) - 1
	minY := min(p1.y, p2.y) + 1
	maxY := max(p1.y, p2.y) - 1

	// Segment bounds
	segMinX := min(segment.p1.x, segment.p2.x)
	segMaxX := max(segment.p1.x, segment.p2.x)
	segMinY := min(segment.p1.y, segment.p2.y)
	segMaxY := max(segment.p1.y, segment.p2.y)

	// Check for non-overlap
	if segMaxX < minX || segMinX > maxX || segMaxY < minY || segMinY > maxY {
		return false
	}
	return true
}

// isInside checks if a point (x,y) is inside the polygon using ray casting
func isInside(x, y float64, polygon []Point) bool {
	inside := false
	n := len(polygon)
	for i := 0; i < n; i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%n]

		// Check if ray intersects edge
		// Ray cast to the right
		if (float64(p1.y) > y) != (float64(p2.y) > y) {
			intersectX := (float64(p2.x)-float64(p1.x))*(y-float64(p1.y))/(float64(p2.y)-float64(p1.y)) + float64(p1.x)
			if x < intersectX {
				inside = !inside
			}
		}
	}
	return inside
}

// min returns the min of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
