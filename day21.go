package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type day21Player struct {
	hitPoints int
	damage    int
	armor     int
}

type day21Item struct {
	cost   int
	damage int
	armor  int
}

type Day21Puzzle struct {
	boss day21Player
}

func NewDay21(lines []string) (Day21Puzzle, error) {
	if len(lines) != 3 {
		return Day21Puzzle{}, fmt.Errorf("invalid input")
	}
	hitPoints, err := day21ParseStat(lines[0], "Hit Points:")
	if err != nil {
		return Day21Puzzle{}, err
	}
	damage, err := day21ParseStat(lines[1], "Damage:")
	if err != nil {
		return Day21Puzzle{}, err
	}
	armor, err := day21ParseStat(lines[2], "Armor:")
	if err != nil {
		return Day21Puzzle{}, err
	}
	return Day21Puzzle{
		boss: day21Player{hitPoints: hitPoints, damage: damage, armor: armor},
	}, nil
}

// Day21 solves day 21 for the selected part.
func Day21(puzzle Day21Puzzle, part1 bool) uint {
	combos := day21ItemCombinations()
	bestWin := int(^uint(0) >> 1)
	bestLose := 0

	for _, gear := range combos {
		me := day21Player{hitPoints: 100, damage: gear.damage, armor: gear.armor}
		if day21Win(me, puzzle.boss) {
			bestWin = min(bestWin, gear.cost)
		} else {
			bestLose = max(bestLose, gear.cost)
		}
	}

	if part1 {
		return uint(bestWin)
	}
	return uint(bestLose)
}

func day21ParseStat(line, prefix string) (int, error) {
	if !strings.HasPrefix(line, prefix) {
		return 0, fmt.Errorf("invalid stat line %q", line)
	}
	n, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, prefix)))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func day21Win(me, boss day21Player) bool {
	for {
		boss.hitPoints -= max(1, me.damage-boss.armor)
		if boss.hitPoints <= 0 {
			return true
		}
		me.hitPoints -= max(1, boss.damage-me.armor)
		if me.hitPoints <= 0 {
			return false
		}
	}
}

var day21Weapons = []day21Item{
	{8, 4, 0},
	{10, 5, 0},
	{25, 6, 0},
	{40, 7, 0},
	{74, 8, 0},
}

var day21Armors = []day21Item{
	{0, 0, 0},
	{13, 0, 1},
	{31, 0, 2},
	{53, 0, 3},
	{75, 0, 4},
	{102, 0, 5},
}

var day21Rings = []day21Item{
	{25, 1, 0},
	{50, 2, 0},
	{100, 3, 0},
	{20, 0, 1},
	{40, 0, 2},
	{80, 0, 3},
}

func day21ItemCombinations() []day21Item {
	m := make(map[day21Item]bool, 1024)
	ringSets := []day21Item{{0, 0, 0}} // no rings
	for i := range len(day21Rings) {
		ringSets = append(ringSets, day21Rings[i]) // one ring
		for j := i + 1; j < len(day21Rings); j++ { // two rings
			ringSets = append(ringSets, day21Item{
				cost:   day21Rings[i].cost + day21Rings[j].cost,
				damage: day21Rings[i].damage + day21Rings[j].damage,
				armor:  day21Rings[i].armor + day21Rings[j].armor,
			})
		}
	}

	for _, weapon := range day21Weapons {
		for _, armor := range day21Armors {
			for _, rings := range ringSets {
					item := day21Item{
						cost:   weapon.cost + armor.cost + rings.cost,
						damage: weapon.damage + armor.damage + rings.damage,
						armor:  weapon.armor + armor.armor + rings.armor,
					}
					m[item] = true
			}
		}
	}
	out := make([]day21Item, 0, len(m))
	for it := range m {
		out = append(out, it)
	}
	return out
}
