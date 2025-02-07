package main

import (
	"fmt"
	"math"
	"time"
	"utils"
)

const day = 21

// Character represents the Boss and The Player as well.
type Character struct {
	hp, dmg, armor int
}

// Weapon is a representation of a weapon with damage and cost
type Weapon struct {
	name   string
	cost   int
	damage int
}

// Armor representation of a armor with armor value
type Armor struct {
	name  string
	cost  int
	armor int
}

// DefenseRing rings which improve armor
type DefenseRing struct {
	name    string
	cost    int
	defense int
}

// DamageRing rings which improve damage
type DamageRing struct {
	name   string
	cost   int
	damage int
}

// Shop represents the shop which has a variety of items
type Shop struct {
	weapons      map[int]Weapon
	armors       map[int]Armor
	defenseRings map[int]DefenseRing
	damageRings  map[int]DamageRing
}

var shop Shop
var itemCombinations []int

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

	smallestCost := math.MaxInt64
	var (
		weapondmg int
		armor     int
		defring   int
		dmgring   int
		cWeapond  int
		cArmor    int
		cDefRing  int
		cDmgRing  int
	)

	for _, v := range shop.weapons {
		weapondmg = v.damage
		cWeapond = v.cost
		for a := 0; a <= len(shop.armors); a++ {
			armor = shop.armors[a].armor
			cArmor = shop.armors[a].cost
			for defr := 0; defr <= len(shop.defenseRings); defr++ {
				defring = shop.defenseRings[defr].defense
				cDefRing = shop.defenseRings[defr].cost

				for dmgr := 0; dmgr <= len(shop.damageRings); dmgr++ {
					dmgring = shop.damageRings[dmgr].damage
					cDmgRing = shop.damageRings[dmgr].cost

					moneySpent := cWeapond + cArmor + cDefRing + cDmgRing
					playersTurn := true

					player := &Character{hp: 100, dmg: weapondmg + dmgring, armor: armor + defring}
					boss := createBoss(data)
					for {
						switch playersTurn {
						case true:
							player.attack(boss)
							playersTurn = false
						case false:
							boss.attack(player)
							playersTurn = true
						}

						if player.hp <= 0 || boss.hp <= 0 {
							break
						}
					}

					if player.hp > 0 {
						if moneySpent < smallestCost {
							smallestCost = moneySpent
						}
					}
				}
			}
		}
	}
	solution = smallestCost
	return solution
}

func solutionB() int {
	var solution = 0
	data, err := utils.ReadFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	maxCost := 0
	var (
		weapondmg int
		armor     int
		defring   int
		dmgring   int
		cWeapond  int
		cArmor    int
		cDefRing  int
		cDmgRing  int
	)

	for _, v := range shop.weapons {
		weapondmg = v.damage
		cWeapond = v.cost
		for a := 0; a <= len(shop.armors); a++ {
			armor = shop.armors[a].armor
			cArmor = shop.armors[a].cost
			for defr := 0; defr <= len(shop.defenseRings); defr++ {
				defring = shop.defenseRings[defr].defense
				cDefRing = shop.defenseRings[defr].cost

				for dmgr := 0; dmgr <= len(shop.damageRings); dmgr++ {
					dmgring = shop.damageRings[dmgr].damage
					cDmgRing = shop.damageRings[dmgr].cost

					moneySpent := cWeapond + cArmor + cDefRing + cDmgRing
					playersTurn := true

					player := &Character{hp: 100, dmg: weapondmg + dmgring, armor: armor + defring}
					boss := createBoss(data)
					for {
						switch playersTurn {
						case true:
							player.attack(boss)
							playersTurn = false
						case false:
							boss.attack(player)
							playersTurn = true
						}

						if player.hp <= 0 || boss.hp <= 0 {
							break
						}
					}

					if boss.hp > 0 {
						if moneySpent > maxCost {
							maxCost = moneySpent
						}
					}
				}
			}
		}
	}
	solution = maxCost
	return solution
}

// attack attacks the other character
func (c1 *Character) attack(c2 *Character) {
	dmg := c1.dmg - c2.armor
	if dmg <= 0 {
		dmg = 1
	}
	c2.hp -= dmg
}

// init initializes the shop with the items
func init() {
	shop = Shop{
		weapons: map[int]Weapon{
			0: {"Dagger", 8, 4},
			1: {"Shortsword", 10, 5},
			2: {"Warhammer", 25, 6},
			3: {"Longsword", 40, 7},
			4: {"Greataxe", 74, 8},
		},
		//Starts from 1 because 0 will mark that it's an optional buy
		armors: map[int]Armor{
			0: {"Nothing", 0, 0},
			1: {"Leather", 13, 1},
			2: {"Chainmail", 31, 2},
			3: {"Splintmail", 53, 3},
			4: {"Bandedmail", 75, 4},
			5: {"Platemail", 102, 5},
		},
		//Starts from 1 because 0 will mark that it's an optional buy
		defenseRings: map[int]DefenseRing{
			0: {"Nothing", 0, 0},
			1: {"Defense +1", 20, 1},
			2: {"Defense +2", 40, 2},
			3: {"Defense +3", 80, 3},
		},
		//Starts from 1 because 0 will mark that it's an optional buy
		damageRings: map[int]DamageRing{
			0: {"Nothing", 0, 0},
			1: {"Damage +1", 25, 1},
			2: {"Damage +2", 50, 2},
			3: {"Damage +3", 100, 3},
		},
	}
}

func createBoss(input string) *Character {
	var hp, dmg, armor int
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d\nArmor: %d", &hp, &dmg, &armor)
	return &Character{hp: hp, dmg: dmg, armor: armor}
}
