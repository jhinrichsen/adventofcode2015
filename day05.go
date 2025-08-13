package adventofcode2015

import (
	"bytes"
	"strings"
)

const (
	nice    = true
	naughty = false
)

// Day5Part1 returns number of nice strings.
func Day5Part1(input []string) (n uint) {
	return day5(input, nicePart1)
}

// Day5Part2 returns number of nice strings.
func Day5Part2(input []string) (n uint) {
	return day5(input, nicePart2)
}

func day5(input []string, f func(string) bool) (n uint) {
	for _, s := range input {
		if f(s) {
			n++
		}
	}
	return
}

func nicePart1(s string) bool {
	return p1(s) && p2(s) && p3(s)
}

func nicePart2(s string) bool {
	return p4(s) && p5(s)
}

// does not contain ab, cd, pq, or xy
func p1(s string) bool {
	s = strings.ToLower(s)
	return !(strings.Contains(s, "ab") ||
		strings.Contains(s, "cd") ||
		strings.Contains(s, "pq") ||
		strings.Contains(s, "xy"))
}

// at least three vowels
func p2(s string) bool {
	var n uint
	for _, r := range s {
		switch r {
		case 'A', 'a':
			n++
		case 'E', 'e':
			n++
		case 'I', 'i':
			n++
		case 'O', 'o':
			n++
		case 'U', 'u':
			n++
		}
	}
	return n >= 3
}

// at least one letter that appears twice in a row.
// Only works for non-unicode strings, i.e. plain ASCII.
func p3(s string) bool {
	buf := []byte(s)
	for i := 1; i < len(buf); i++ {
		if buf[i-1] == buf[i] {
			return nice
		}
	}
	return naughty
}

// a pair of any two letters that appears at least twice in the string without
// overlapping.
// Only works for non-unicode strings, i.e. plain ASCII.
func p4(s string) bool {
	if len(s) < 4 {
		return naughty
	}
	buf := []byte(s)
	for i := 0; i < len(buf)-3; i++ {
		pair := buf[i : i+2]
		right := buf[i+2 : len(s)]
		if bytes.Contains(right, pair) {
			return nice
		}
	}
	return naughty
}

// at least one letter which repeats with exactly one letter between them.
// Only works for non-unicode strings, i.e. plain ASCII.
func p5(s string) bool {
	if len(s) < 4 {
		return false
	}
	buf := []byte(s)
	for i := 1; i < len(buf)-1; i++ {
		if buf[i-1] == buf[i+1] {
			return nice
		}
	}
	return naughty
}
