package adventofcode2015

import (
	"fmt"
	"sort"
	"strings"
)

const day19Separator = "=>"

// Plant holds a day 19 Red-Nosed Reindeer nuclear fusion/fission plant.
type plant struct {
	replacements map[string][]string
	molecule     string
}

// newPlant creates a plant from a list of replacements such as "H => OH".
func newPlant(replacements []string) (plant, error) {
	var p plant
	p.replacements = make(map[string][]string, len(replacements))
	for i, s := range replacements {
		parts := strings.Split(s, day19Separator)
		if len(parts) != 2 {
			return p, fmt.Errorf("illegal format at index %d: "+
				"want H => OH but got %q", i, s)
		}
		p.addReplacement(strings.TrimSpace(parts[0]),
			strings.TrimSpace(parts[1]))
	}
	return p, nil
}

func (a plant) addReplacement(k, v string) {
	if existing, ok := a.replacements[k]; ok {
		// add to existing replacement
		a.replacements[k] = append(existing, v)
	} else {
		vs := make([]string, 1)
		vs[0] = v
		a.replacements[k] = vs
	}
}

func (a plant) distinct() map[string]bool {
	m := make(map[string]bool)
	for from, v := range a.replacements {
		for _, into := range v {
			for c := strings.Count(a.molecule, from); c > 0; c-- {
				r := replaceNth(a.molecule, from, into, c)
				m[r] = true
			}
		}
	}
	return m
}

// Day19Part1 returns number of possible medicine molecules.
func Day19Part1(p plant) uint {
	return uint(len(p.distinct()))
}

type reducer struct {
	from, into string
}

// reducers returns the complement of replacers, if O => OH is a replacer,
// OH => O is a reducer.
func (a plant) reducers() (rs []reducer) {
	for k, vs := range a.replacements {
		for _, v := range vs {
			rs = append(rs, reducer{v, k})
		}

	}
	return rs
}

func reduced0(m map[string]bool) bool {
	return m["e"]
}

// Day19Part2 returns number of steps to convert 'e' to plant's molecule.
func Day19Part2(molecule string, rs []reducer) (step uint) {
	m := make(map[string]bool)
	m[molecule] = true
	return reduce(m, rs)
}

func reduce(prospects map[string]bool, rs []reducer) uint {
	for step := uint(1); ; step++ {
		reduced := make(map[string]bool)
		for p := range prospects {
			for _, r := range rs {
				for c := strings.Count(p, r.from); c >= 0; c-- {
					s := replaceNth(p, r.from, r.into, c)
					reduced[s] = true
				}
			}
		}

		// end result available in prospects?
		if reduced0(reduced) {
			return step
		}
		if len(reduced) == 0 {
			return 0
		}

		// heuristically reduce complexity: build a list of shortest prospects
		var ss []string
		for k := range reduced {
			ss = append(ss, k)
		}
		sort.Sort(ByLen(ss))

		// consider shortest only
		next := make(map[string]bool)
		fittest := len(ss)
		const heuristic = 100
		if fittest > heuristic {
			fittest = heuristic
		}
		for _, s := range ss[:fittest] {
			next[s] = true
		}
		prospects = next
	}
}

// ByLen sorts strings by length of string.
type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// replaceNth replaces the nth occurrence of old in s with new.
func replaceNth(s, old, new string, n int) string {
	i := 0
	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
