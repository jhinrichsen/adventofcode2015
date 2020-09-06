package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type happiness struct {
	attendee  string
	neighbour string
	change    int
}

type happinesses []happiness

func (as happinesses) attendees() attendees {
	m := make(map[string]bool)
	for _, h := range as {
		m[h.attendee] = true
		m[h.neighbour] = true
	}
	return keys(m)
}

type attendees []string

// index returns index of a in as, or error if not found.
func (as attendees) index(a string) (int, error) {
	for i := range as {
		if as[i] == a {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

// left treats list of attendees as ring and returns previous index, optionally
// wrapping.
func (as attendees) left(idx int) int {
	if idx == 0 {
		return len(as) - 1
	}
	return idx - 1
}

func (as attendees) right(idx int) int {
	if idx == len(as)-1 {
		return 0
	}
	return idx + 1
}

func (as attendees) sitsNext(att1, att2 string) bool {
	idx1, err := as.index(att1)
	if err != nil {
		return false
	}
	if as[as.left(idx1)] == att2 || as[as.right(idx1)] == att2 {
		return true
	}
	return false
}

func (as happinesses) change(atts attendees) int {
	var n int
	for _, a := range as {
		if atts.sitsNext(a.attendee, a.neighbour) {
			n += a.change
		}
	}
	return n
}

// Day13 returns total change in happiness.
func Day13(filename string) (int, error) {
	var max int
	hs, err := newHappinesses(filename)
	if err != nil {
		return max, err
	}
	perms := make(chan []string)
	atts := hs.attendees()
	go heap(len(atts), atts, perms)
	for perm := range perms {
		n := hs.change(perm)
		if err != nil {
			return max, err
		}
		if n > max {
			max = n
		}
	}
	return max, nil
}

// newHappinesses parses lines in the form "Alice would gain 54 happiness units by
// sitting next to Bob."
func newHappinesses(filename string) (happinesses, error) {
	var hs []happiness
	lines, err := linesFromFilename(filename)
	if err != nil {
		return hs, err
	}
	for i, line := range lines {
		parts := strings.Fields(strings.TrimSuffix(line, "."))
		n, err := strconv.Atoi(parts[3])
		if err != nil {
			return hs, fmt.Errorf("error parsing line %d: %v", i, err)
		}
		if parts[2] == "lose" {
			n = -n
		}
		hs = append(hs, happiness{
			parts[0],
			parts[10],
			n,
		})
	}
	return hs, nil
}
