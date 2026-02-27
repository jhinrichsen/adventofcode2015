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
	buf    []byte
	width  int
	height int
}

// newGrid parses lines of #. combinations, separated by newline.
func newGrid(s string) (grid, error) {
	var g grid
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		line := sc.Bytes()
		if g.width == 0 {
			g.width = len(line)
		}
		g.buf = append(g.buf, line...)
		g.height++
	}
	return g, nil
}

func (a grid) dim() (uint, uint) {
	return uint(a.width), uint(a.height)
}

// on returns number of lights in grid that are "on".
func (a grid) on() uint {
	var n uint
	for _, b := range a.buf {
		if b == lightOn {
			n++
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
	w, h := a.width, a.height
	a.buf[0] = b             // top-left
	a.buf[w-1] = b           // top-right
	a.buf[(h-1)*w] = b       // bottom-left
	a.buf[(h-1)*w+(w-1)] = b // bottom-right
}

func (a *grid) step() {
	next := make([]byte, len(a.buf))
	g := Grid{W: a.width, H: a.height}

	for idx, nbrs := range g.C8Indices() {
		current := a.buf[idx]

		var count uint
		for nidx := range nbrs {
			if a.buf[nidx] == lightOn {
				count++
			}
		}

		next[idx] = nextStatus(current == lightOn, count)
	}
	a.buf = next
}

func (a grid) String() string {
	var sb strings.Builder
	for y := 0; y < a.height; y++ {
		sb.Write(a.buf[y*a.width : (y+1)*a.width])
		if y < a.height-1 {
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
