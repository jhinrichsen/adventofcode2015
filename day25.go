package adventofcode2015

import (
	"fmt"
	"math"
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

func (a day25State) x() uint {
	m := uint(math.Floor((-1 + math.Sqrt(float64(8*a.n-7))) / 2))
	return a.n - m*(m+1)/2
}

func (a day25State) y() uint {
	t := uint(math.Floor(-1+math.Sqrt(float64(8*a.n-7)))) / 2
	return (t*t+3*t+4)/2 - a.n
}

func day25CodeAt(x, y uint) uint {
	d := newDay25State()
	for !(d.x() == x && d.y() == y) {
		d.next()
	}
	return d.code
}
