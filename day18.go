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
	next := make([]byte, len(buf))

	if !part1 {
		day18SetCorners(buf, puzzle.w, puzzle.h, day18LightOn)
	}
	for range day18Steps {
		day18StepInto(buf, next, puzzle.w, puzzle.h)
		buf, next = next, buf
		if !part1 {
			day18SetCorners(buf, puzzle.w, puzzle.h, day18LightOn)
		}
	}
	return day18CountOn(buf)
}

func day18Step(buf []byte, w, h int) []byte {
	next := make([]byte, len(buf))
	day18StepInto(buf, next, w, h)
	return next
}

func day18StepInto(buf, next []byte, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x
			n := uint(0)
			y0 := max(0, y-1)
			y1 := min(h-1, y+1)
			x0 := max(0, x-1)
			x1 := min(w-1, x+1)
			for yy := y0; yy <= y1; yy++ {
				row := yy * w
				for xx := x0; xx <= x1; xx++ {
					if xx == x && yy == y {
						continue
					}
					if buf[row+xx] == day18LightOn {
						n++
					}
				}
			}
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
