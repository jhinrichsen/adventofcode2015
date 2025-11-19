package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type sizes [3]uint

func newSizes(s string) (sizes, error) {
	var ss sizes
	parts := strings.Split(s, "x")
	if len(ss) != len(parts) {
		return ss, fmt.Errorf("want %d parts but got %d", len(ss), len(parts))
	}
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return ss, fmt.Errorf("error parsing %q: part %d is not a number", s, i)
		}
		ss[i] = uint(n)
	}
	return ss, nil
}

// Day02Part1 returns sum of sizes.
func Day02Part1(lines []string) (uint, error) {
	var sum uint
	for _, line := range lines {
		s, err := newSizes(line)
		if err != nil {
			return sum, err
		}
		sum += s.size()
	}
	return sum, nil
}

func (a sizes) size() uint {
	s1 := a[0] * a[1]
	s2 := a[1] * a[2]
	s3 := a[2] * a[0]
	return 2*(s1+s2+s3) + min(s1, min(s2, s3))
}

// Day02Part2 returns the length of the ribbon band.
func Day02Part2(lines []string) (uint, error) {
	var sum uint
	for _, line := range lines {
		s, err := newSizes(line)
		if err != nil {
			return sum, err
		}
		sum += s.ribbon()
	}
	return sum, nil
}

func (a sizes) ribbon() uint {
	return a.bow() + a.present()
}

func (a sizes) bow() uint {
	return a[0] * a[1] * a[2]
}

func (a sizes) present() uint {
	s1 := 2 * (a[0] + a[1])
	s2 := 2 * (a[1] + a[2])
	s3 := 2 * (a[2] + a[0])
	return min(s1, min(s2, s3))
}
