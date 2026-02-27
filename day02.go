package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type sizes [3]uint
type Day02Puzzle []sizes

func NewDay02(lines []string) (Day02Puzzle, error) {
	puzzle := make(Day02Puzzle, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "x")
		var s sizes
		if len(s) != len(parts) {
			return nil, fmt.Errorf("want %d parts but got %d", len(s), len(parts))
		}
		for i, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				return nil, fmt.Errorf("error parsing %q: part %d is not a number", line, i)
			}
			s[i] = uint(n)
		}
		puzzle = append(puzzle, s)
	}
	return puzzle, nil
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
