package adventofcode2015

import (
	"fmt"
)

type sizes [3]uint
type Day02Puzzle []sizes

func NewDay02(lines []string) (Day02Puzzle, error) {
	puzzle := make(Day02Puzzle, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		s, err := parseDay02Sizes(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing %q: %w", line, err)
		}
		puzzle = append(puzzle, s)
	}
	return puzzle, nil
}

func parseDay02Sizes(line string) (sizes, error) {
	var out sizes
	part := 0
	n := uint(0)
	hasDigit := false
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			n = n*10 + uint(c-'0')
			hasDigit = true
			continue
		}
		if c == 'x' {
			if !hasDigit || part >= len(out)-1 {
				return sizes{}, fmt.Errorf("invalid separator at %d", i)
			}
			out[part] = n
			part++
			n = 0
			hasDigit = false
			continue
		}
		return sizes{}, fmt.Errorf("invalid character %q", c)
	}
	if part != len(out)-1 || !hasDigit {
		return sizes{}, fmt.Errorf("invalid dimensions")
	}
	out[part] = n
	return out, nil
}

// Day02 solves day 2 for the selected part.
func Day02(puzzle Day02Puzzle, part1 bool) uint {
	var sum uint
	for _, s := range puzzle {
		if part1 {
			sum += wrappingPaperSize(s)
		} else {
			sum += ribbonLength(s)
		}
	}
	return sum
}

func wrappingPaperSize(s sizes) uint {
	s1 := s[0] * s[1]
	s2 := s[1] * s[2]
	s3 := s[2] * s[0]
	return 2*(s1+s2+s3) + min(s1, min(s2, s3))
}

func ribbonLength(s sizes) uint {
	return bowLength(s) + smallestPerimeter(s)
}

func bowLength(s sizes) uint {
	return s[0] * s[1] * s[2]
}

func smallestPerimeter(s sizes) uint {
	s1 := 2 * (s[0] + s[1])
	s2 := 2 * (s[1] + s[2])
	s3 := 2 * (s[2] + s[0])
	return min(s1, min(s2, s3))
}
