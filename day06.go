package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	width  = 1000
	height = 1000
)

type light byte
type lightMutator func(light) light

type lights [width][height]light

type day06Op uint8

const (
	day06On day06Op = iota
	day06Off
	day06Toggle
)

type day06Instruction struct {
	op     day06Op
	x1, y1 uint
	x2, y2 uint
}

type Day06Puzzle []day06Instruction

func newLights() lights {
	var l [width][height]light
	return l
}

func (a *lights) mut(x1, y1, x2, y2 uint, f lightMutator) {
	// all coordinates are inclusive
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			a[x][y] = f(a[x][y])
		}
	}
}

func (a lights) count() (n uint) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			n += uint(a[x][y])
		}
	}
	return
}

// parseCoord converts "500,500" into (500,500).
func parseCoord(s string) (uint, uint, error) {
	ps := strings.Split(s, ",")
	if len(ps) != 2 {
		return 0, 0, fmt.Errorf("want two comma separated numbers but got %q", s)
	}
	x, err := strconv.Atoi(ps[0])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse x: %q", ps[0])
	}
	y, err := strconv.Atoi(ps[1])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot parse y: %q", ps[1])
	}
	return uint(x), uint(y), nil
}

func NewDay06(lines []string) (Day06Puzzle, error) {
	puzzle := make(Day06Puzzle, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		var op day06Op
		var firstField, secondField int
		if parts[0] == "toggle" {
			op = day06Toggle
			firstField = 1
			secondField = 3
		} else {
			if parts[1] == "on" {
				op = day06On
			} else {
				op = day06Off
			}
			firstField = 2
			secondField = 4
		}
		x1, y1, err := parseCoord(parts[firstField])
		if err != nil {
			return nil, err
		}
		x2, y2, err := parseCoord(parts[secondField])
		if err != nil {
			return nil, err
		}
		puzzle = append(puzzle, day06Instruction{op: op, x1: x1, y1: y1, x2: x2, y2: y2})
	}
	return puzzle, nil
}

// Day06 solves day 6 for the selected part.
func Day06(puzzle Day06Puzzle, part1 bool) uint {
	muts := [3]lightMutator{}
	if part1 {
		muts = [...]lightMutator{
			func(_ light) light { return 1 },
			func(_ light) light { return 0 },
			func(b light) light { return 1 - b },
		}
	} else {
		muts = [...]lightMutator{
			func(b light) light { return b + 1 },
			func(b light) light {
				if b == 0 {
					return 0
				}
				return b - 1
			},
			func(b light) light { return b + 2 },
		}
	}

	ls := newLights()
	for _, ins := range puzzle {
		ls.mut(ins.x1, ins.y1, ins.x2, ins.y2, muts[ins.op])
	}
	return ls.count()
}
