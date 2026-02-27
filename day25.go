package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type day25State struct {
	n    uint
	code uint
}

type Day25Puzzle struct {
	row uint
	col uint
}

func NewDay25(lines []string) (Day25Puzzle, error) {
	if len(lines) != 1 {
		return Day25Puzzle{}, fmt.Errorf("invalid input")
	}
	fields := strings.Fields(lines[0])
	if len(fields) < 18 {
		return Day25Puzzle{}, fmt.Errorf("invalid input line")
	}
	row, err := strconv.ParseUint(strings.TrimSuffix(fields[15], ","), 10, 64)
	if err != nil {
		return Day25Puzzle{}, err
	}
	col, err := strconv.ParseUint(strings.TrimSuffix(fields[17], "."), 10, 64)
	if err != nil {
		return Day25Puzzle{}, err
	}
	return Day25Puzzle{row: uint(row), col: uint(col)}, nil
}

// Day25 solves day 25.
func Day25(puzzle Day25Puzzle, _ bool) uint {
	return day25CodeAt(puzzle.col, puzzle.row)
}

func newDay25State() day25State {
	return day25State{n: 1, code: 20151125}
}

func (a *day25State) next() {
	a.n++
	a.code = a.code * 252533 % 33554393
}

func day25CodeAt(x, y uint) uint {
	const (
		start = uint(20151125)
		mul   = uint(252533)
		mod   = uint(33554393)
	)
	diag := x + y - 1
	n := (diag-1)*diag/2 + x
	return start * powMod(mul, n-1, mod) % mod
}

func powMod(base, exp, mod uint) uint {
	result := uint(1)
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}
