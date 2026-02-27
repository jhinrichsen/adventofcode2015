package adventofcode2015

import (
	"fmt"
)

const (
	width  = 1000
	height = 1000
)

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

func NewDay06(lines []string) (Day06Puzzle, error) {
	puzzle := make(Day06Puzzle, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		var op day06Op
		idx := 0
		if len(line) >= 7 && line[:7] == "toggle " {
			op = day06Toggle
			idx = 7
		} else if len(line) >= 8 && line[:8] == "turn on " {
			op = day06On
			idx = 8
		} else if len(line) >= 9 && line[:9] == "turn off " {
			op = day06Off
			idx = 9
		} else {
			return nil, fmt.Errorf("invalid instruction %q", line)
		}

		x1, y1, next, ok := day06ParseCoord(line, idx)
		if !ok {
			return nil, fmt.Errorf("invalid first coordinate in %q", line)
		}
		const through = " through "
		if next+len(through) > len(line) || line[next:next+len(through)] != through {
			return nil, fmt.Errorf("missing through in %q", line)
		}
		x2, y2, end, ok := day06ParseCoord(line, next+len(through))
		if !ok || end != len(line) {
			return nil, fmt.Errorf("invalid second coordinate in %q", line)
		}
		puzzle = append(puzzle, day06Instruction{op: op, x1: x1, y1: y1, x2: x2, y2: y2})
	}
	return puzzle, nil
}

func day06ParseCoord(s string, i int) (x, y uint, next int, ok bool) {
	xv, i, ok := day06ParseUint(s, i)
	if !ok || i >= len(s) || s[i] != ',' {
		return 0, 0, 0, false
	}
	yv, i, ok := day06ParseUint(s, i+1)
	if !ok {
		return 0, 0, 0, false
	}
	return xv, yv, i, true
}

func day06ParseUint(s string, i int) (uint, int, bool) {
	if i >= len(s) || s[i] < '0' || s[i] > '9' {
		return 0, 0, false
	}
	n := uint(0)
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		n = n*10 + uint(s[i]-'0')
		i++
	}
	return n, i, true
}

// Day06 solves day 6 for the selected part.
func Day06(puzzle Day06Puzzle, part1 bool) uint {
	if part1 {
		var grid [width][height]byte
		for _, ins := range puzzle {
			switch ins.op {
			case day06On:
				for x := ins.x1; x <= ins.x2; x++ {
					for y := ins.y1; y <= ins.y2; y++ {
						grid[x][y] = 1
					}
				}
			case day06Off:
				for x := ins.x1; x <= ins.x2; x++ {
					for y := ins.y1; y <= ins.y2; y++ {
						grid[x][y] = 0
					}
				}
			case day06Toggle:
				for x := ins.x1; x <= ins.x2; x++ {
					for y := ins.y1; y <= ins.y2; y++ {
						grid[x][y] ^= 1
					}
				}
			}
		}
		total := uint(0)
		for x := range width {
			for y := range height {
				total += uint(grid[x][y])
			}
		}
		return total
	}

	var grid [width][height]uint16
	for _, ins := range puzzle {
		switch ins.op {
		case day06On:
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					grid[x][y]++
				}
			}
		case day06Off:
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					if grid[x][y] > 0 {
						grid[x][y]--
					}
				}
			}
		case day06Toggle:
			for x := ins.x1; x <= ins.x2; x++ {
				for y := ins.y1; y <= ins.y2; y++ {
					grid[x][y] += 2
				}
			}
		}
	}
	total := uint(0)
	for x := range width {
		for y := range height {
			total += uint(grid[x][y])
		}
	}
	return total
}
