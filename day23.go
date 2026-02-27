package adventofcode2015

import (
	"strconv"
	"strings"
)

type Day23Puzzle []string

func NewDay23(lines []string) (Day23Puzzle, error) {
	puzzle := make(Day23Puzzle, len(lines))
	for i := range lines {
		puzzle[i] = strings.ReplaceAll(lines[i], ",", "")
	}
	return puzzle, nil
}

// Day23 solves day 23 for the selected part.
func Day23(puzzle Day23Puzzle, part1 bool) uint {
	initialA := uint(0)
	if !part1 {
		initialA = 1
	}
	_, b := day23Run(puzzle, initialA, 0)
	return b
}

func day23Run(instructions []string, initialA, initialB uint) (a uint, b uint) {
	a = initialA
	b = initialB

	register := func(reg string) *uint {
		if reg == "a" {
			return &a
		}
		if reg == "b" {
			return &b
		}
		return nil
	}
	offset := func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return i
	}

	for pc := 0; pc >= 0 && pc < len(instructions); {
		fs := strings.Fields(instructions[pc])
		if len(fs) < 2 {
			break
		}
		op1 := fs[1]
		switch fs[0] {
		case "hlf":
			r := register(op1)
			if r == nil {
				return
			}
			*r /= 2
			pc++
		case "tpl":
			r := register(op1)
			if r == nil {
				return
			}
			*r *= 3
			pc++
		case "inc":
			r := register(op1)
			if r == nil {
				return
			}
			*r = *r + 1
			pc++
		case "jmp":
			pc += offset(op1)
		case "jie":
			if len(fs) < 3 {
				return
			}
			r := register(op1)
			if r == nil {
				return
			}
			if *r%2 == 0 {
				pc += offset(fs[2])
			} else {
				pc++
			}
		case "jio":
			if len(fs) < 3 {
				return
			}
			r := register(op1)
			if r == nil {
				return
			}
			if *r == 1 {
				pc += offset(fs[2])
			} else {
				pc++
			}
		default:
			return
		}
	}
	return
}
