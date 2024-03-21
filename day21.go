package adventofcode2015

import (
	"sort"
)

type day21Player struct {
	hitPoints int  // also known as life, score
	items     item // all the gear he is carrying (one or more)
}

func (a day21Player) alive() bool {
	return a.hitPoints > 0
}

var day21Boss = day21Player{
	hitPoints: 100,
	items: item{
		damage: 8,
		armor:  2,
	},
}
var nakedMe = day21Player{100, item{0, 0, 0}}

// duel plays until one player dies.
func duel(players *[2]day21Player) {
	attack := 0 // player 1 begins
	turn := func() {
		attack = 1 - attack
	}
	for {
		damage := players[attack].items.damage -
			players[1-attack].items.armor
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
	cost, damage, armor int
}

func (a *item) add(i item) {
	a.cost += i.cost
	a.damage += i.damage
	a.armor += i.armor
}

// Must always have one weapon, even in part 2.
var weaponItems = []item{
	{8, 4, 0},  // Dagger
	{10, 5, 0}, // Shortsword
	{25, 6, 0}, // Warhammer
	{40, 7, 0}, // Longsword
	{74, 8, 0}, // Greataxe
}

var armorItems = []item{
	{0, 0, 0},   // No armor
	{13, 0, 1},  // Leather
	{31, 0, 2},  // Chainmail
	{53, 0, 3},  // Splitmail
	{75, 0, 4},  // Bandedmail
	{102, 0, 5}, // Platemail
}

var ringItems = []item{
	{0, 0, 0},   // No ring
	{25, 1, 0},  // Damage +1
	{50, 2, 0},  // Damage +2
	{100, 3, 0}, // Damage +3
	{20, 0, 1},  // Defense +1
	{40, 0, 2},  // Defense +2
	{80, 0, 3},  // Defense +3
}

// itemCombinations returns array sorted by cost asc.
// i got part 2 wrong, it seems as if one cannot carry any combination, but
// still need exactly one weapon.
func itemCombinations() []item {
	m := make(map[item]bool, 5*6*15)
	// 1 weapon
	for _, weapon := range weaponItems {
		for _, armor := range armorItems {
			for _, ring1 := range ringItems {
				for _, ring2 := range ringItems {
					// cannot use the same ring twice
					if ring1 == ring2 {
						continue
					}
					w := weapon
					w.add(armor)
					w.add(ring1)
					w.add(ring2)
					m[w] = true
				}
			}
		}
	}
	// map -> array
	var cs []item
	for k := range m {
		cs = append(cs, k)
	}
	return cs
}

// Day21Part1 returns minimal cost to survive.
func Day21Part1() int {
	cs := itemCombinations()
	// sort by cost asc
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].cost < cs[j].cost
	})
	for _, c := range cs {
		p := nakedMe
		p.items.armor = c.armor
		p.items.damage = c.damage
		players := [...]day21Player{
			p,
			day21Boss,
		}
		duel(&players)
		if players[0].alive() {
			return c.cost
		}
	}
	return 0
}

// Day21Part2 returns highest cost so that player 0 still loses.
func Day21Part2() int {
	cs := itemCombinations()
	// sort by cost desc
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].cost > cs[j].cost
	})
	for _, items := range cs {
		p := nakedMe
		p.items = items
		players := [...]day21Player{
			p,
			day21Boss,
		}
		duel(&players)
		if !players[0].alive() {
			return players[0].items.cost
		}
	}
	return 0
}
