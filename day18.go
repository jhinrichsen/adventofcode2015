package adventofcode2015

import (
	"bufio"
	"strings"
)

const (
	lightOn  = '#'
	lightOff = '.'
)

type grid struct {
	buf [][]byte
}

// newGrid parses lines of #. combinations, separated by newline.
func newGrid(s string) (grid, error) {
	var g grid
	g.buf = make([][]byte, 0)
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		bs := sc.Bytes()
		buf := make([]byte, len(bs))
		copy(buf, bs)
		g.buf = append(g.buf, buf)
	}
	return g, nil
}

func (a grid) dim() (uint, uint) {
	return uint(len(a.buf[0])), uint(len(a.buf))
}

// on returns number of lights in grid that are "on".
func (a grid) on() uint {
	var n uint
	for y := range a.buf {
		for x := range a.buf[0] {
			if a.buf[y][x] == lightOn {
				n++
			}
		}
	}
	return n
}

func (a *grid) setCorners(t bool) {
	var b byte
	if t {
		b = lightOn
	} else {
		b = lightOff
	}
	a.buf[0][0] = b
	a.buf[0][len(a.buf[0])-1] = b
	a.buf[len(a.buf)-1][0] = b
	a.buf[len(a.buf)-1][len(a.buf[0])-1] = b
}

func (a *grid) step() {
	// render into offline frame
	dimy := len(a.buf)
	dimx := len(a.buf[0])
	buf := make([][]byte, dimy)
	var n uint
	check := func(x, y int) {
		if 0 <= y &&
			y < dimy &&
			0 <= x &&
			x < dimx &&
			a.buf[y][x] == lightOn {

			n++
		}
	}
	for y := range a.buf {
		buf[y] = make([]byte, dimx)
		for x := range a.buf[0] {
			n = 0
			check(x, y-1)   // N
			check(x+1, y-1) // NE
			check(x+1, y)   // E
			check(x+1, y+1) // SE
			check(x, y+1)   // S
			check(x-1, y+1) // SW
			check(x-1, y)   // W
			check(x-1, y-1) // NW

			buf[y][x] = nextStatus(a.buf[y][x] == lightOn, n)
		}
	}
	a.buf = buf
}

func (a grid) String() string {
	var sb strings.Builder
	for y := 0; y < len(a.buf); y++ {
		sb.Write(a.buf[y])
		// No newline for last element
		if y < len(a.buf)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

// next returns status of next iteration, depending on current status and number
// of neighbours that are on.
func nextStatus(isOn bool, n uint) byte {
	if isOn {
		if n == 2 || n == 3 {
			return lightOn
		}
		return lightOff
	}
	if n == 3 {
		return lightOn
	}
	return lightOff
}

// Day18Part1 returns number of lights that are on after n steps.
// Conway's Game of life.
func Day18Part1(g grid, steps uint) uint {
	for ; steps > 0; steps-- {
		g.step()
	}
	return g.on()
}

// Day18Part2 returns number of lights that are on after n steps under the
// precondition that all corner lights are always on.
func Day18Part2(g grid, steps uint) uint {
	for ; steps > 0; steps-- {
		g.setCorners(true)
		g.step()
	}
	g.setCorners(true)
	return g.on()
}
