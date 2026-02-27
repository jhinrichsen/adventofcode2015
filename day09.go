package adventofcode2015

import (
	"strconv"
	"strings"
)

type Day09Puzzle struct {
	dist [][]uint
	has  [][]bool
}

func NewDay09(lines []string) (Day09Puzzle, error) {
	cityID := make(map[string]int)
	type edge struct {
		a int
		b int
		d uint
	}
	edges := make([]edge, 0, len(lines))

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 5 {
			return Day09Puzzle{}, strconv.ErrSyntax
		}
		d, err := strconv.ParseUint(parts[4], 10, 64)
		if err != nil {
			return Day09Puzzle{}, err
		}

		a, ok := cityID[parts[0]]
		if !ok {
			a = len(cityID)
			cityID[parts[0]] = a
		}
		b, ok := cityID[parts[2]]
		if !ok {
			b = len(cityID)
			cityID[parts[2]] = b
		}
		edges = append(edges, edge{a: a, b: b, d: uint(d)})
	}

	n := len(cityID)
	dist := make([][]uint, n)
	has := make([][]bool, n)
	for y := range n {
		dist[y] = make([]uint, n)
		has[y] = make([]bool, n)
	}
	for _, e := range edges {
		dist[e.a][e.b] = e.d
		dist[e.b][e.a] = e.d
		has[e.a][e.b] = true
		has[e.b][e.a] = true
	}

	return Day09Puzzle{dist: dist, has: has}, nil
}

func Day09(puzzle Day09Puzzle, part1 bool) uint {
	n := len(puzzle.dist)
	if n == 0 {
		return 0
	}

	ids := make([]int, n)
	for i := range n {
		ids[i] = i
	}
	found := false

	eval := func() (uint, bool) {
		var sum uint
		for i := 1; i < n; i++ {
			a := ids[i-1]
			b := ids[i]
			if !puzzle.has[a][b] {
				return 0, false
			}
			sum += puzzle.dist[a][b]
		}
		return sum, true
	}

	best := uint(0)
	updateBest := func(sum uint) {
		if !found {
			best = sum
			found = true
			return
		}
		if part1 {
			best = min(best, sum)
		} else {
			best = max(best, sum)
		}
	}

	if sum, ok := eval(); ok {
		updateBest(sum)
	}

	// Iterative Heap permutation traversal, evaluated in-place.
	c := make([]int, n)
	for i := 0; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				ids[0], ids[i] = ids[i], ids[0]
			} else {
				ids[c[i]], ids[i] = ids[i], ids[c[i]]
			}
			if sum, ok := eval(); ok {
				updateBest(sum)
			}
			c[i]++
			i = 0
			continue
		}
		c[i] = 0
		i++
	}

	if !found {
		return 0
	}
	return best
}

func keys(m map[string]bool) []string {
	ss := make([]string, 0, len(m))
	for k := range m {
		ss = append(ss, k)
	}
	return ss
}
