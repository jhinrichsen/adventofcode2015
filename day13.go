package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type Day13Puzzle struct {
	change [][]int
}

func NewDay13(lines []string) (Day13Puzzle, error) {
	type edge struct {
		a int
		b int
		n int
	}

	ids := make(map[string]int)
	edges := make([]edge, 0, len(lines))
	for i, line := range lines {
		parts := strings.Fields(strings.TrimSuffix(line, "."))
		if len(parts) != 11 {
			return Day13Puzzle{}, fmt.Errorf("invalid line %d", i+1)
		}
		n, err := strconv.Atoi(parts[3])
		if err != nil {
			return Day13Puzzle{}, fmt.Errorf("line %d: %w", i+1, err)
		}
		if parts[2] == "lose" {
			n = -n
		}

		a, ok := ids[parts[0]]
		if !ok {
			a = len(ids)
			ids[parts[0]] = a
		}
		b, ok := ids[parts[10]]
		if !ok {
			b = len(ids)
			ids[parts[10]] = b
		}
		edges = append(edges, edge{a: a, b: b, n: n})
	}

	size := len(ids)
	change := make([][]int, size)
	for y := range size {
		change[y] = make([]int, size)
	}
	for _, e := range edges {
		change[e.a][e.b] = e.n
	}
	return Day13Puzzle{change: change}, nil
}

// Day13 solves day 13 for the selected part.
func Day13(puzzle Day13Puzzle, part1 bool) uint {
	n := len(puzzle.change)
	if n == 0 {
		return 0
	}
	size := n
	if !part1 {
		size++
	}

	ids := make([]uint, size)
	for i := range size {
		ids[i] = uint(i)
	}

	perms := make(chan []uint)
	go heapUint(size, ids, perms)

	best := 0
	found := false
	for perm := range perms {
		total := 0
		for i := range size {
			a := int(perm[i])
			b := int(perm[(i+1)%size])
			if a < n && b < n {
				total += puzzle.change[a][b] + puzzle.change[b][a]
			}
		}
		if !found || total > best {
			best = total
			found = true
		}
	}
	if !found || best < 0 {
		return 0
	}
	return uint(best)
}

