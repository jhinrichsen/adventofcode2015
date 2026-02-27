package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	day15Calories = 500
	day15TSP      = 100
)

type day15Ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Day15Puzzle []day15Ingredient

func NewDay15(lines []string) (Day15Puzzle, error) {
	puzzle := make(Day15Puzzle, 0, len(lines))
	for i, line := range lines {
		ing, err := day15ParseIngredient(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}
		puzzle = append(puzzle, ing)
	}
	return puzzle, nil
}

// Day15 solves day 15 for the selected part.
func Day15(puzzle Day15Puzzle, part1 bool) uint {
	n := len(puzzle)
	if n == 0 {
		return 0
	}
	best := uint(0)
	amounts := make([]int, n)
	remaining := make([]int, n)
	next := make([]int, n)
	depth := 0
	remaining[0] = day15TSP
	for depth >= 0 {
		if depth == n-1 {
			amounts[depth] = remaining[depth]
			best = max(best, day15ScoreWithCalories(puzzle, amounts, part1))
			depth--
			if depth >= 0 {
				next[depth]++
			}
			continue
		}
		if next[depth] > remaining[depth] {
			next[depth] = 0
			depth--
			if depth >= 0 {
				next[depth]++
			}
			continue
		}
		amounts[depth] = next[depth]
		depth++
		remaining[depth] = remaining[depth-1] - amounts[depth-1]
		next[depth] = 0
	}
	return best
}

func day15ParseIngredient(s string) (day15Ingredient, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return day15Ingredient{}, fmt.Errorf("invalid ingredient")
	}
	props := strings.Split(parts[1], ",")
	if len(props) != 5 {
		return day15Ingredient{}, fmt.Errorf("invalid properties")
	}
	vals := [5]int{}
	for i, prop := range props {
		fields := strings.Fields(prop)
		if len(fields) != 2 {
			return day15Ingredient{}, fmt.Errorf("invalid property %q", prop)
		}
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			return day15Ingredient{}, err
		}
		vals[i] = n
	}
	return day15Ingredient{
		capacity:   vals[0],
		durability: vals[1],
		flavor:     vals[2],
		texture:    vals[3],
		calories:   vals[4],
	}, nil
}

func day15ScoreWithCalories(puzzle Day15Puzzle, amounts []int, part1 bool) uint {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0
	for i, ing := range puzzle {
		a := amounts[i]
		capacity += a * ing.capacity
		durability += a * ing.durability
		flavor += a * ing.flavor
		texture += a * ing.texture
		calories += a * ing.calories
	}
	if !part1 && calories != day15Calories {
		return 0
	}
	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}
	return uint(capacity * durability * flavor * texture)
}
