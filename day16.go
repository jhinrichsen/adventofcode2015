package adventofcode2015

import (
	"fmt"
)

type day16Sue struct {
	number uint
	props  [10]uint
	has    [10]bool
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
	for _, sue := range puzzle {
		if part1 {
			if day16MatchExact(sue) {
				return sue.number
			}
			continue
		}
		if day16MatchRanges(sue) {
			return sue.number
		}
	}
	return 0
}

func day16ParseSue(s string) (day16Sue, error) {
	if len(s) < 6 || s[0] != 'S' || s[1] != 'u' || s[2] != 'e' || s[3] != ' ' {
		return day16Sue{}, fmt.Errorf("invalid line")
	}
	i := 4
	n := uint(0)
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		n = n*10 + uint(s[i]-'0')
		i++
	}
	if i >= len(s) || s[i] != ':' {
		return day16Sue{}, fmt.Errorf("invalid sue id")
	}
	i++
	if i < len(s) && s[i] == ' ' {
		i++
	}

	sue := day16Sue{number: uint(n)}
	for i < len(s) {
		start := i
		for i < len(s) && s[i] != ':' {
			i++
		}
		if i >= len(s) {
			return day16Sue{}, fmt.Errorf("invalid property")
		}
		id, ok := day16PropIDBytes(s[start:i])
		if !ok {
			return day16Sue{}, fmt.Errorf("invalid property")
		}
		i++
		if i < len(s) && s[i] == ' ' {
			i++
		}
		v := uint(0)
		if i >= len(s) || s[i] < '0' || s[i] > '9' {
			return day16Sue{}, fmt.Errorf("invalid property value")
		}
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			v = v*10 + uint(s[i]-'0')
			i++
		}
		sue.props[id] = v
		sue.has[id] = true
		if i >= len(s) {
			break
		}
		if s[i] != ',' {
			return day16Sue{}, fmt.Errorf("invalid property separator")
		}
		i++
		if i < len(s) && s[i] == ' ' {
			i++
		}
	}
	return sue, nil
}

const (
	day16Children = iota
	day16Cats
	day16Samoyeds
	day16Pomeranians
	day16Akitas
	day16Vizslas
	day16Goldfish
	day16Trees
	day16Cars
	day16Perfumes
)

var day16Target = [10]uint{3, 7, 2, 3, 0, 0, 5, 3, 2, 1}

func day16PropID(name string) (int, bool) {
	switch name {
	case "children":
		return day16Children, true
	case "cats":
		return day16Cats, true
	case "samoyeds":
		return day16Samoyeds, true
	case "pomeranians":
		return day16Pomeranians, true
	case "akitas":
		return day16Akitas, true
	case "vizslas":
		return day16Vizslas, true
	case "goldfish":
		return day16Goldfish, true
	case "trees":
		return day16Trees, true
	case "cars":
		return day16Cars, true
	case "perfumes":
		return day16Perfumes, true
	default:
		return -1, false
	}
}

func day16PropIDBytes(name string) (int, bool) {
	switch len(name) {
	case 4:
		if name == "cats" {
			return day16Cats, true
		}
		if name == "cars" {
			return day16Cars, true
		}
	case 5:
		if name == "trees" {
			return day16Trees, true
		}
	case 6:
		if name == "akitas" {
			return day16Akitas, true
		}
	case 7:
		if name == "vizslas" {
			return day16Vizslas, true
		}
	case 8:
		if name == "goldfish" {
			return day16Goldfish, true
		}
		if name == "children" {
			return day16Children, true
		}
		if name == "samoyeds" {
			return day16Samoyeds, true
		}
		if name == "perfumes" {
			return day16Perfumes, true
		}
	case 11:
		if name == "pomeranians" {
			return day16Pomeranians, true
		}
	}
	return -1, false
}

func day16MatchExact(sue day16Sue) bool {
	for i := range sue.has {
		if sue.has[i] && sue.props[i] != day16Target[i] {
			return false
		}
	}
	return true
}

func day16MatchRanges(sue day16Sue) bool {
	for i := range sue.has {
		if !sue.has[i] {
			continue
		}
		v := sue.props[i]
		want := day16Target[i]
		switch i {
		case day16Cats, day16Trees:
			if v <= want {
				return false
			}
		case day16Pomeranians, day16Goldfish:
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
