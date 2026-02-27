package adventofcode2015

import (
	"fmt"
	"strconv"
)

type Day20Puzzle struct {
	target uint
}

func NewDay20(lines []string) (Day20Puzzle, error) {
	if len(lines) != 1 {
		return Day20Puzzle{}, fmt.Errorf("invalid input")
	}
	n, err := strconv.ParseUint(lines[0], 10, 64)
	if err != nil {
		return Day20Puzzle{}, err
	}
	return Day20Puzzle{target: uint(n)}, nil
}

// Day20 solves day 20 for the selected part.
func Day20(puzzle Day20Puzzle, part1 bool) uint {
	if part1 {
		return day20Part1(puzzle.target)
	}
	return day20Part2(puzzle.target)
}

func day20Part1(target uint) uint {
	houses := make([]uint, target/10+1)
	for elf := uint(1); elf < uint(len(houses)); elf++ {
		for house := elf; house < uint(len(houses)); house += elf {
			houses[house] += elf * 10
		}
	}
	for house := uint(1); house < uint(len(houses)); house++ {
		if houses[house] >= target {
			return house
		}
	}
	return 0
}

func day20Part2(target uint) uint {
	houses := make([]uint, target/11+1)
	for elf := uint(1); elf < uint(len(houses)); elf++ {
		limit := min(uint(len(houses)-1), elf*50)
		for house := elf; house <= limit; house += elf {
			houses[house] += elf * 11
		}
	}
	for house := uint(1); house < uint(len(houses)); house++ {
		if houses[house] >= target {
			return house
		}
	}
	return 0
}

func day20PresentsPart1(house uint) uint {
	if house == 0 {
		return 0
	}
	sum := uint(0)
	for elf := uint(1); elf <= house; elf++ {
		if house%elf == 0 {
			sum += elf * 10
		}
	}
	return sum
}

