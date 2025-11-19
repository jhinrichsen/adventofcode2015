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

// convert "500,500" into (500,500)
func atot(s string) (uint, uint, error) {
	ps := strings.Split(s, ",")
	if len(ps) != 2 {
		return 0, 0, fmt.Errorf("want two comma separated numbers but got %q",
			s)
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

// Day06Part1 returns number of lights in state 'on'.
func Day06Part1(lines []string) (uint, error) {
	return day6(lines, [...]lightMutator{
		// On
		func(b light) light {
			return 1
		},
		// Off
		func(b light) light {
			return 0
		},
		func(b light) light {
			return 1 - b
		}})
}

// Day06Part2 returns number of lights in state 'on'.
func Day06Part2(lines []string) (uint, error) {
	return day6(lines, [...]lightMutator{
		// On
		func(b light) light {
			return b + 1
		},
		// Off
		func(b light) light {
			if b == 0 {
				return 0
			}
			return b - 1
		},
		func(b light) light {
			return b + 2
		}})
}

// mutator order: on, off, toggle
func day6(lines []string, muts [3]lightMutator) (uint, error) {
	ls := newLights()
	for _, line := range lines {
		var f lightMutator
		// indices for (x,y) fields into commands
		var firstField, secondField int
		parts := strings.Fields(line)
		if parts[0] == "toggle" {
			f = muts[2]
			firstField = 1
			secondField = 3
		} else {
			if parts[1] == "on" {
				f = muts[0]
			} else {
				f = muts[1]
			}
			firstField = 2
			secondField = 4
		}
		x1, y1, err := atot(parts[firstField])
		if err != nil {
			return 0, err
		}
		x2, y2, err := atot(parts[secondField])
		if err != nil {
			return 0, err
		}
		ls.mut(x1, y1, x2, y2, f)
	}
	return ls.count(), nil

}
