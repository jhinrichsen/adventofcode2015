package adventofcode2015

import (
	"sort"
)

type player struct {
	hitPoints     int
	damage, armor uint
}

func (a player) alive() bool {
	return a.hitPoints > 0
}

// Day21Part1 plays until one player dies.
func Day21Part1(players *[2]player) {
	attack := 0 // player 1 begins
	turn := func() {
		attack = 1 - attack
	}
	for {
		damage := players[attack].damage - players[1-attack].armor
		if damage < 1 {
			damage = 1
		}
		players[1-attack].hitPoints -= int(damage)
		if !players[1-attack].alive() {
			return
		}
		turn()
	}
}

type item struct {
	cost, damage, armor uint
}

func (a *item) add(i item) {
	a.cost += i.cost
	a.damage += i.damage
	a.armor += i.armor
}

var weaponItems = []item{
	{8, 4, 0},  // Dagger
	{10, 5, 0}, // Shortsword
	{25, 6, 0}, // Warhammer
	{40, 7, 0}, // Longsword
	{74, 8, 0}, // Greataxe
}

var armorItems = []item{
	{13, 0, 1},  // Leather
	{31, 0, 2},  // Chainmail
	{53, 0, 3},  // Splitmail
	{75, 0, 4},  // Bandedmail
	{102, 0, 5}, // Platemail
}

var ringItems = []item{
	{25, 1, 0},  // Damage +1
	{50, 2, 0},  // Damage +2
	{100, 3, 0}, // Damage +3
	{20, 0, 1},  // Defense +1
	{40, 0, 2},  // Defense +2
	{80, 0, 3},  // Defense +3
}

// combinations returns array sorted by cost asc.
func combinations() []item {
	m := make(map[item]bool, 5*6*15)
	// 1 weapon
	for _, weapon := range weaponItems {
		// 0..1 armor
		for a := 0; a <= 1; a++ {
			for _, armor := range armorItems {
				// 0..2 rings
				for r := 0; r <= 2; r++ {
					for _, ring1 := range ringItems {
						for _, ring2 := range ringItems {
							w := weapon
							if a == 1 {
								w.add(armor)
							}
							if r == 1 {
								w.add(ring1)
							}
							// cannot use the same ring twice
							if r == 2 && (ring1 != ring2) {
								w.add(ring2)
							}
							m[w] = true
						}
					}
				}
			}
		}
	}
	// map -> array
	var cs []item
	for k := range m {
		cs = append(cs, k)
	}
	// sort by cost asc
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].cost < cs[j].cost
	})
	return cs
}
