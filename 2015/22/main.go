package main

import (
	"fmt"
	"math"
	"time"
	"utils"
)

const (
	day             = 22
	playerStartHP   = 50
	playerStartMana = 500
	//bossStartHP     = 71
	//bossDamage      = 10
)

// GameState is the game state structure
type GameState struct {
	playerHP, playerMana, bossHP, bossDamage, manaSpent int
	shieldTimer, poisonTimer, rechargeTimer             int
	playerTurn, isHardMode                              bool
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

	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	bossHP, bossDamage := getBossStatsFromInput(data)
	solution = run(bossHP, bossDamage, false)
	return solution
}

func solutionB() int {
	var solution = 0

	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	bossHP, bossDamage := getBossStatsFromInput(data)
	solution = run(bossHP, bossDamage, true)
	return solution
}

// getBossStatsFromInput parses the input string and returns the boss's hit points and damage
func getBossStatsFromInput(input string) (int, int) {
	var hp, dmg int
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d", &hp, &dmg)
	return hp, dmg
}

// applyEffects applies active effects
func applyEffects(state *GameState) {
	if state.shieldTimer > 0 {
		state.shieldTimer--
	}
	if state.poisonTimer > 0 {
		state.bossHP -= 3
		state.poisonTimer--
	}
	if state.rechargeTimer > 0 {
		state.playerMana += 101
		state.rechargeTimer--
	}
}

// playerTurn simulates the player's turn
func playerTurn(state GameState, minManaSpent *int) {
	// in Hard, Player loses 1 hit point at the start of their turn
	if state.isHardMode {
		state.playerHP--
		if state.playerHP <= 0 {
			return
		}
	}

	applyEffects(&state)

	if state.bossHP <= 0 {
		if state.manaSpent < *minManaSpent {
			*minManaSpent = state.manaSpent
		}
		return
	}

	if state.playerMana < 53 {
		return
	}

	state.playerTurn = false

	// Magic Missile
	if state.playerMana >= 53 {
		newState := state
		newState.playerMana -= 53
		newState.manaSpent += 53
		newState.bossHP -= 4
		bossTurn(newState, minManaSpent)
	}

	// Drain
	if state.playerMana >= 73 {
		newState := state
		newState.playerMana -= 73
		newState.manaSpent += 73
		newState.bossHP -= 2
		newState.playerHP += 2
		bossTurn(newState, minManaSpent)
	}

	// Shield
	if state.playerMana >= 113 && state.shieldTimer == 0 {
		newState := state
		newState.playerMana -= 113
		newState.manaSpent += 113
		newState.shieldTimer = 6
		bossTurn(newState, minManaSpent)
	}

	// Poison
	if state.playerMana >= 173 && state.poisonTimer == 0 {
		newState := state
		newState.playerMana -= 173
		newState.manaSpent += 173
		newState.poisonTimer = 6
		bossTurn(newState, minManaSpent)
	}

	// Recharge
	if state.playerMana >= 229 && state.rechargeTimer == 0 {
		newState := state
		newState.playerMana -= 229
		newState.manaSpent += 229
		newState.rechargeTimer = 5
		bossTurn(newState, minManaSpent)
	}
}

// bossTurn simulates the boss's turn
func bossTurn(state GameState, minManaSpent *int) {
	applyEffects(&state)

	if state.bossHP <= 0 {
		if state.manaSpent < *minManaSpent {
			*minManaSpent = state.manaSpent
		}
		return
	}

	armor := 0
	if state.shieldTimer > 0 {
		armor = 7
	}

	damage := state.bossDamage - armor
	if damage < 1 {
		damage = 1
	}

	state.playerHP -= damage

	if state.playerHP > 0 {
		state.playerTurn = true
		playerTurn(state, minManaSpent)
	}
}

// Main function to initialize game state and start the simulation
func run(bossStartHP, bossDamage int, isHardMode bool) int {
	initialState := GameState{
		playerHP:      playerStartHP,
		playerMana:    playerStartMana,
		bossHP:        bossStartHP,
		bossDamage:    bossDamage,
		manaSpent:     0,
		shieldTimer:   0,
		poisonTimer:   0,
		rechargeTimer: 0,
		playerTurn:    true,
		isHardMode:    isHardMode,
	}

	minManaSpent := math.MaxInt32
	playerTurn(initialState, &minManaSpent)
	//fmt.Println("Minimum mana spent to win:", minManaSpent)
	return minManaSpent
}
