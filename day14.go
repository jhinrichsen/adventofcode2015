package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

// Reindeer holds a day 15 domain model.
type Reindeer struct {
	Name     string
	Velocity uint
	Flying   uint
	Resting  uint
}

func max(a, b uint) uint {
	if a < b {
		return b
	}
	return a
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// km returns distance in km after n seconds.
func (a Reindeer) km(sec uint) uint {
	oneFrame := a.Flying + a.Resting

	// full frames
	frames := sec / oneFrame
	kmPerFrame := a.Flying * a.Velocity // [s] * [km/s]
	d1 := frames * kmPerFrame           // [] * [km]

	// partial frames
	// only seconds in flight mode will add to distance
	d2 := min(sec%oneFrame, a.Flying) * a.Velocity // [s] * [km/s]
	return d1 + d2
}

func newReindeer(line string) (Reindeer, error) {
	fields := strings.Fields(line)
	if len(fields) != 15 {
		return Reindeer{},
			fmt.Errorf("want 15 fields but got %d", len(fields))
	}
	const (
		base = 10
		bits = 8 // examples and puzzle input have very small values
	)
	i3, err := strconv.ParseUint(fields[3], base, bits)
	if err != nil {
		return Reindeer{},
			fmt.Errorf("col %d: no number: %s", 3, fields[3])
	}
	i6, err := strconv.ParseUint(fields[6], base, bits)
	if err != nil {
		return Reindeer{},
			fmt.Errorf("col %d: no number: %s", 6, fields[6])
	}
	i13, err := strconv.ParseUint(fields[13], base, bits)
	if err != nil {
		return Reindeer{},
			fmt.Errorf("col %d: no number: %s", 13, fields[13])
	}

	return Reindeer{
		fields[0],
		uint(i3),
		uint(i6),
		uint(i13),
	}, nil
}

// Day14 returns maximum distance in km after sec seconds.
func Day14(rs []Reindeer, sec uint) uint {
	var d uint
	for _, r := range rs {
		d = max(d, r.km(sec))
	}
	return d
}
