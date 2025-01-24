package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

const day = 15

const maxTeaspoons = 100

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

type Recipe struct {
	Amounts []int
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

	ingredients := parseLines(lines)

	solution = run(ingredients, false)

	return solution
}

func run(ingredients []Ingredient, checkCalories bool) int {
	var maxScore = 0

	recipe := Recipe{
		Amounts: make([]int, len(ingredients)),
	}

	for i := 0; i <= maxTeaspoons; i++ {
		recipe.Amounts[0] = i
		for j := 0; j <= maxTeaspoons-i; j++ {
			recipe.Amounts[1] = j
			for k := 0; k <= maxTeaspoons-i-j; k++ {
				recipe.Amounts[2] = k
				// The remaining amount must go to the last ingredient
				recipe.Amounts[3] = maxTeaspoons - i - j - k

				// Calculate score
				if score, valid := getScore(recipe, ingredients, checkCalories); valid && score > maxScore {
					maxScore = score
				}
			}
		}
	}
	return maxScore
}

func parseLines(lines []string) []Ingredient {
	var ingredients []Ingredient

	for _, line := range lines {
		var name string
		var capacity, durability, flavor, texture, calories int

		_, err := fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name, &capacity, &durability, &flavor, &texture, &calories)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		in := Ingredient{
			Name:       strings.TrimSuffix(name, ":"),
			Capacity:   capacity,
			Durability: durability,
			Flavor:     flavor,
			Texture:    texture,
			Calories:   calories,
		}
		ingredients = append(ingredients, in)
	}
	return ingredients
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	ingredients := parseLines(lines)

	solution = run(ingredients, true)

	return solution
}

// getScore calculates the score of a recipe
func getScore(recipe Recipe, ingredients []Ingredient, checkCalories bool) (int, bool) {
	var capacity, durability, flavor, texture, calories int

	for i, in := range ingredients {
		amount := recipe.Amounts[i]

		capacity += in.Capacity * amount
		durability += in.Durability * amount
		flavor += in.Flavor * amount
		texture += in.Texture * amount
		calories += in.Calories * amount
	}

	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0, false
	}

	// If we need to check calories (Part B), we return 0 if they are not 500
	if checkCalories && calories != 500 {
		return 0, false
	}

	return capacity * durability * flavor * texture, true
}
