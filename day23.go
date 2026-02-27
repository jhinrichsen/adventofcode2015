package adventofcode2015

import (
	"strconv"
	"strings"
)

type day23Op uint8

const (
	day23Hlf day23Op = iota
	day23Tpl
	day23Inc
	day23Jmp
	day23Jie
	day23Jio
)

type day23Instruction struct {
	op  day23Op
	reg byte
	off int
}

type Day23Puzzle []day23Instruction

func NewDay23(lines []string) (Day23Puzzle, error) {
	puzzle := make(Day23Puzzle, 0, len(lines))
	for _, line := range lines {
		ins, ok := day23ParseInstruction(line)
		if !ok {
			continue
		}
		puzzle = append(puzzle, ins)
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

func day23Run(instructions Day23Puzzle, initialA, initialB uint) (a uint, b uint) {
	a = initialA
	b = initialB

	reg := func(r byte) *uint {
		if r == 'a' {
			return &a
		}
		if r == 'b' {
			return &b
		}
		return nil
	}

	for pc := 0; pc >= 0 && pc < len(instructions); {
		ins := instructions[pc]
		switch ins.op {
		case day23Hlf:
			r := reg(ins.reg)
			if r == nil {
				return
			}
			*r /= 2
			pc++
		case day23Tpl:
			r := reg(ins.reg)
			if r == nil {
				return
			}
			*r *= 3
			pc++
		case day23Inc:
			r := reg(ins.reg)
			if r == nil {
				return
			}
			*r++
			pc++
		case day23Jmp:
			pc += ins.off
		case day23Jie:
			r := reg(ins.reg)
			if r == nil {
				return
			}
			if *r%2 == 0 {
				pc += ins.off
			} else {
				pc++
			}
		case day23Jio:
			r := reg(ins.reg)
			if r == nil {
				return
			}
			if *r == 1 {
				pc += ins.off
			} else {
				pc++
			}
		default:
			return
		}
	}
	return
}

func day23ParseInstruction(line string) (day23Instruction, bool) {
	line = strings.ReplaceAll(line, ",", "")
	fs := strings.Fields(line)
	if len(fs) < 2 {
		return day23Instruction{}, false
	}
	atoi := func(s string) int {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return n
	}
	switch fs[0] {
	case "hlf":
		return day23Instruction{op: day23Hlf, reg: fs[1][0]}, true
	case "tpl":
		return day23Instruction{op: day23Tpl, reg: fs[1][0]}, true
	case "inc":
		return day23Instruction{op: day23Inc, reg: fs[1][0]}, true
	case "jmp":
		return day23Instruction{op: day23Jmp, off: atoi(fs[1])}, true
	case "jie":
		if len(fs) < 3 {
			return day23Instruction{}, false
		}
		return day23Instruction{op: day23Jie, reg: fs[1][0], off: atoi(fs[2])}, true
	case "jio":
		if len(fs) < 3 {
			return day23Instruction{}, false
		}
		return day23Instruction{op: day23Jio, reg: fs[1][0], off: atoi(fs[2])}, true
	default:
		return day23Instruction{}, false
	}
}
