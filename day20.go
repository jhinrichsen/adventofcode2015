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
	n := int(target/10 + 1)
	houses := make([]uint32, n)
	for elf := 1; elf < n; elf++ {
		presents := uint32(elf * 10)
		for house := elf; house < n; house += elf {
			houses[house] += presents
		}
	}
	target32 := uint32(target)
	for house := 1; house < n; house++ {
		if houses[house] >= target32 {
			return uint(house)
		}
	}
	return 0
}

func day20Part2(target uint) uint {
	n := int(target/11 + 1)
	houses := make([]uint32, n)
	for elf := 1; elf < n; elf++ {
		limit := min(n-1, elf*50)
		presents := uint32(elf * 11)
		for house := elf; house <= limit; house += elf {
			houses[house] += presents
		}
	}
	target32 := uint32(target)
	for house := 1; house < n; house++ {
		if houses[house] >= target32 {
			return uint(house)
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
