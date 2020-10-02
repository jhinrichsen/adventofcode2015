package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

// Day23Part1 runs instructions and returns register a and b.
// Implemented purely on string, no mnemonic e.a. (interpreter vs. compiler)
func Day23Part1(instructions []string) (a uint, b uint) {
	return day23(instructions, 0, 0)
}

// Day23Part2 runs instructions and returns register a and b.
// Implemented purely on string, no mnemonic e.a. (interpreter vs. compiler)
func Day23Part2(instructions []string, initialA uint) (a uint, b uint) {
	return day23(instructions, initialA, 0)
}

func day23(instructions []string, initialA, initialB uint) (a uint, b uint) {
	// remove any "," so we get clean 1/2/3 fields
	for i := range instructions {
		instructions[i] = strings.ReplaceAll(instructions[i], ",", "")
	}

	a = initialA
	b = initialB

	register := func(reg string) *uint {
		if reg == "a" {
			return &a
		} else if reg == "b" {
			return &b
		}
		panic(fmt.Sprintf("unknown register %q", reg))
	}
	hlf := func(reg *uint) {
		*reg = *reg / 2
	}
	tpl := func(reg *uint) {
		*reg = *reg * 3
	}
	inc := func(reg *uint) {
		*reg = *reg + 1
	}
	offset := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	isEven := func(n uint) bool {
		return n%2 == 0
	}
	isOne := func(n uint) bool {
		return n == 1
	}

	var op1, op2 string
	for pc := 0; pc < len(instructions); {
		fs := strings.Fields(instructions[pc])
		op1 = fs[1]
		switch fs[0] {
		case "hlf":
			hlf(register(op1))
			pc++
		case "tpl":
			tpl(register(op1))
			pc++
		case "inc":
			inc(register(op1))
			pc++
		case "jmp":
			pc += offset(op1)
		case "jie":
			op2 = fs[2]
			if isEven(*register(op1)) {
				pc += offset(op2)
			} else {
				pc++
			}
		case "jio":
			op2 = fs[2]
			if isOne(*register(op1)) {
				pc += offset(op2)
			} else {
				pc++
			}
		default:
			return
		}
	}
	return
}
