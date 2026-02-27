package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type day16Sue struct {
	number uint
	props  map[string]uint
}

type Day16Puzzle []day16Sue

func NewDay16(lines []string) (Day16Puzzle, error) {
	puzzle := make(Day16Puzzle, 0, len(lines))
	for i, line := range lines {
		sue, err := day16ParseSue(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}
		puzzle = append(puzzle, sue)
	}
	return puzzle, nil
}

// Day16 solves day 16 for the selected part.
func Day16(puzzle Day16Puzzle, part1 bool) uint {
	target := day16Target()
	for _, sue := range puzzle {
		if part1 {
			if day16MatchExact(target, sue.props) {
				return sue.number
			}
			continue
		}
		if day16MatchRanges(target, sue.props) {
			return sue.number
		}
	}
	return 0
}

func day16ParseSue(s string) (day16Sue, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return day16Sue{}, fmt.Errorf("invalid line")
	}
	idFields := strings.Fields(parts[0])
	if len(idFields) != 2 {
		return day16Sue{}, fmt.Errorf("invalid sue id")
	}
	n, err := strconv.ParseUint(idFields[1], 10, 64)
	if err != nil {
		return day16Sue{}, err
	}

	props := strings.Split(parts[1], ",")
	m := make(map[string]uint, len(props))
	for _, prop := range props {
		ps := strings.Split(prop, ":")
		if len(ps) != 2 {
			return day16Sue{}, fmt.Errorf("invalid property %q", prop)
		}
		v, err := strconv.ParseUint(strings.TrimSpace(ps[1]), 10, 64)
		if err != nil {
			return day16Sue{}, err
		}
		m[strings.TrimSpace(ps[0])] = uint(v)
	}
	return day16Sue{number: uint(n), props: m}, nil
}

func day16Target() map[string]uint {
	return map[string]uint{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

func day16MatchExact(target, props map[string]uint) bool {
	for k, v := range props {
		if want, ok := target[k]; ok && v != want {
			return false
		}
	}
	return true
}

func day16MatchRanges(target, props map[string]uint) bool {
	for k, v := range props {
		want, ok := target[k]
		if !ok {
			continue
		}
		switch k {
		case "cats", "trees":
			if v <= want {
				return false
			}
		case "pomeranians", "goldfish":
			if v >= want {
				return false
			}
		default:
			if v != want {
				return false
			}
		}
	}
	return true
}

