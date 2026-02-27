package adventofcode2015

import "strings"

// Day05 solves day 5 for the selected part.
func Day05(input []string, part1 bool) (n uint) {
	rule := nicePart2
	if part1 {
		rule = nicePart1
	}
	for _, s := range input {
		if rule(s) {
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
			return true
		}
	}
	return false
}

// a pair of any two letters that appears at least twice in the string without
// overlapping.
// Only works for non-unicode strings, i.e. plain ASCII.
func p4(s string) bool {
	if len(s) < 4 {
		return false
	}
	buf := []byte(s)
	for i := 0; i < len(buf)-1; i++ {
		if buf[i] < 'a' || buf[i] > 'z' || buf[i+1] < 'a' || buf[i+1] > 'z' {
			return p4Generic(buf)
		}
	}

	// seen stores first index+1 where a pair appeared.
	// 0 means unseen; checking (i+1)-seen[pair] >= 2 guarantees non-overlap.
	var seen [26 * 26]int
	for i := 0; i < len(buf)-1; i++ {
		pair := int(buf[i]-'a')*26 + int(buf[i+1]-'a')
		if seen[pair] != 0 && (i+1)-seen[pair] >= 2 {
			return true
		}
		if seen[pair] == 0 {
			seen[pair] = i + 1
		}
	}
	return false
}

func p4Generic(buf []byte) bool {
	for i := 0; i < len(buf)-3; i++ {
		a0, a1 := buf[i], buf[i+1]
		for j := i + 2; j < len(buf)-1; j++ {
			if a0 == buf[j] && a1 == buf[j+1] {
				return true
			}
		}
	}
	return false
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
			return true
		}
	}
	return false
}
