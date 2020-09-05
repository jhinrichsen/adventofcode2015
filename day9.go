package adventofcode2015

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// a->b == b->a
type route struct {
	a        string
	b        string
	distance uint
}

type routes []route

func (a routes) cities() []string {
	m := make(map[string]bool)
	for _, r := range a {
		m[r.a] = true
		m[r.b] = true
	}
	return keys(m)
}

func keys(m map[string]bool) []string {
	var ss []string
	for k := range m {
		ss = append(ss, k)
	}
	return ss
}

func (a routes) distance(cities []string) (uint, error) {
	var n uint
	for i := 1; i < len(cities); i++ {
		found := false
		for _, r := range a {
			if cities[i-1] == r.a && cities[i] == r.b ||
				cities[i-1] == r.b && cities[i] == r.a {

				n += r.distance
				found = true
				break
			}
		}
		if !found {
			return n, fmt.Errorf("no distance for %s -> %s",
				cities[i-1], cities[i])
		}
	}
	return n, nil
}

func newRoutes(filename string) (routes, error) {
	var rs routes
	lines, err := linesFromFilename(filename)
	if err != nil {
		return rs, err
	}
	for _, line := range lines {
		// "London to Dublin = 464"
		parts := strings.Fields(line)
		d, err := strconv.ParseUint(parts[4], 10, 32)
		if err != nil {
			return rs, err
		}
		rs = append(rs, route{
			a:        parts[0],
			b:        parts[2],
			distance: uint(d),
		})
	}
	return rs, nil
}

// Day9 implements the travelling salesman problem, except that Santa Claus does
// not need to end where he started.
// When it comes to generic functions such as permutations(), Go clearly lacks
// generics.
// There's basically three algorithms that fit: Fisher-Yates random shuffle,
// Steinhaus–Johnson–Trotter, and Heap.
func Day9(filename string) (uint, uint, error) {
	min := uint(math.MaxUint32)
	var max uint

	r, err := newRoutes(filename)
	if err != nil {
		return min, max, err
	}
	perms := make(chan []string)
	cs := r.cities()
	go heap(len(cs), cs, perms)
	for perm := range perms {
		n, err := r.distance(perm)
		if err != nil {
			return min, max, err
		}
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return min, max, nil
}
