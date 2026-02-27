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
	if len(puzzle) == 0 {
		return 0
	}

	ch := make(chan []int)
	go KCompositions(day15TSP, len(puzzle), ch)

	best := uint(0)
	for amounts := range ch {
		if !part1 && day15TotalCalories(puzzle, amounts) != day15Calories {
			continue
		}
		best = max(best, day15Score(puzzle, amounts))
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

func day15Score(puzzle Day15Puzzle, amounts []int) uint {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	for i, ing := range puzzle {
		a := amounts[i]
		capacity += a * ing.capacity
		durability += a * ing.durability
		flavor += a * ing.flavor
		texture += a * ing.texture
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

func day15TotalCalories(puzzle Day15Puzzle, amounts []int) int {
	total := 0
	for i, ing := range puzzle {
		total += amounts[i] * ing.calories
	}
	return total
}

