package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type reindeer struct {
	name     string
	velocity uint
	flying   uint
	resting  uint
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
func (a reindeer) km(sec uint) uint {
	oneFrame := a.flying + a.resting

	// full frames
	frames := sec / oneFrame
	kmPerFrame := a.velocity * a.flying // [km/s] * [s]
	d1 := frames * kmPerFrame           // [] * [km]

	// partial frames
	// only seconds in flight mode will add to distance
	d2 := min(sec%oneFrame, a.flying) * a.velocity // [s] * [km/s]
	return d1 + d2
}

func newReindeer(line string) (reindeer, error) {
	fields := strings.Fields(line)
	if len(fields) != 15 {
		return reindeer{},
			fmt.Errorf("want 15 fields but got %d", len(fields))
	}
	const (
		base = 10
		bits = 8 // examples and puzzle input have very small values
	)
	i3, err := strconv.ParseUint(fields[3], base, bits)
	if err != nil {
		return reindeer{},
			fmt.Errorf("col %d: no number: %s", 3, fields[3])
	}
	i6, err := strconv.ParseUint(fields[6], base, bits)
	if err != nil {
		return reindeer{},
			fmt.Errorf("col %d: no number: %s", 6, fields[6])
	}
	i13, err := strconv.ParseUint(fields[13], base, bits)
	if err != nil {
		return reindeer{},
			fmt.Errorf("col %d: no number: %s", 13, fields[13])
	}

	return reindeer{
		fields[0],
		uint(i3),
		uint(i6),
		uint(i13),
	}, nil
}
