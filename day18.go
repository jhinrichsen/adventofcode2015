package adventofcode2015

import "errors"

const (
	day18LightOn  = '#'
	day18LightOff = '.'
	day18Steps    = 100
)

type Day18Puzzle struct {
	buf []byte
	w   int
	h   int
}

func NewDay18(lines []string) (Day18Puzzle, error) {
	if len(lines) == 0 {
		return Day18Puzzle{}, errors.New("empty input")
	}
	w := len(lines[0])
	if w == 0 {
		return Day18Puzzle{}, errors.New("empty line")
	}

	buf := make([]byte, 0, w*len(lines))
	for _, line := range lines {
		if len(line) != w {
			return Day18Puzzle{}, errors.New("non-rectangular grid")
		}
		for i := range len(line) {
			if line[i] != day18LightOn && line[i] != day18LightOff {
				return Day18Puzzle{}, errors.New("invalid char")
			}
		}
		buf = append(buf, []byte(line)...)
	}
	return Day18Puzzle{buf: buf, w: w, h: len(lines)}, nil
}

// Day18 solves day 18 for the selected part.
func Day18(puzzle Day18Puzzle, part1 bool) uint {
	buf := make([]byte, len(puzzle.buf))
	copy(buf, puzzle.buf)

	if !part1 {
		day18SetCorners(buf, puzzle.w, puzzle.h, day18LightOn)
	}
	for range day18Steps {
		buf = day18Step(buf, puzzle.w, puzzle.h)
		if !part1 {
			day18SetCorners(buf, puzzle.w, puzzle.h, day18LightOn)
		}
	}
	return day18CountOn(buf)
}

func day18Step(buf []byte, w, h int) []byte {
	next := make([]byte, len(buf))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x
			n := day18NeighborsOn(buf, w, h, x, y)
			if buf[idx] == day18LightOn {
				if n == 2 || n == 3 {
					next[idx] = day18LightOn
				} else {
					next[idx] = day18LightOff
				}
				continue
			}
			if n == 3 {
				next[idx] = day18LightOn
			} else {
				next[idx] = day18LightOff
			}
		}
	}
	return next
}

func day18NeighborsOn(buf []byte, w, h, x, y int) uint {
	var n uint
	for dy := -1; dy <= 1; dy++ {
		yy := y + dy
		if yy < 0 || yy >= h {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			xx := x + dx
			if xx < 0 || xx >= w {
				continue
			}
			if buf[yy*w+xx] == day18LightOn {
				n++
			}
		}
	}
	return n
}

func day18SetCorners(buf []byte, w, h int, state byte) {
	buf[0] = state
	buf[w-1] = state
	buf[(h-1)*w] = state
	buf[(h-1)*w+(w-1)] = state
}

func day18CountOn(buf []byte) uint {
	var n uint
	for _, b := range buf {
		if b == day18LightOn {
			n++
		}
	}
	return n
}

