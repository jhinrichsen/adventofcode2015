package adventofcode2015

import (
	"strconv"
	"strings"
)

// Sue is identified by a number, and has some properties.
type Sue struct {
	Number     uint
	Properties map[string]uint
}

// parseSue converts a line "Sue 475: trees: 2, cars: 7, akitas: 8" into a map.
func newSue(s string) (Sue, error) {
	parts := strings.SplitN(s, ":", 2)
	n, err := strconv.ParseUint(strings.Fields(parts[0])[1], 10, 32)
	if err != nil {
		return Sue{}, err
	}
	props := strings.Split(parts[1], ",")
	m := make(map[string]uint, len(props))
	for _, prop := range props {
		ps := strings.Split(prop, ":")
		n, err := strconv.ParseUint(strings.TrimSpace(ps[1]), 10, 32)
		if err != nil {
			return Sue{}, err
		}
		m[strings.TrimSpace(ps[0])] = uint(n)
	}
	return Sue{uint(n), m}, nil
}

// match map properties. cannot use reflect.DeepEqual because maps are
// incomplete, so match values on existing keys in both maps.
func match(m1, m2 map[string]uint) bool {
	for k, v := range m1 {
		if v2, ok := m2[k]; ok {
			if v != v2 {
				return false
			}
		}
	}
	return true
}

func matchWorkaroundForBrokenTurboEncapulator(m1, m2 map[string]uint) bool {
	for k, v := range m1 {
		if v2, ok := m2[k]; ok {
			// property dependant equality checks
			switch k {
			case "cats":
				fallthrough
			case "tree":
				if v > v2 {
					return false
				}
			case "pomeranians":
				fallthrough
			case "goldfish":
				if v < v2 {
					return false
				}
			default:
				if v != v2 {
					return false
				}
			}
		}
	}
	return true
}
