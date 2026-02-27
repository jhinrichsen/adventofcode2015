package adventofcode2015

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

const day19Separator = "=>"

type day19Reducer struct {
	from string
	into string
}

type Day19Puzzle struct {
	replacements map[string][]string
	molecule     string
}

func NewDay19(lines []string) (Day19Puzzle, error) {
	if len(lines) < 3 {
		return Day19Puzzle{}, fmt.Errorf("invalid input")
	}
	p := Day19Puzzle{
		replacements: make(map[string][]string, len(lines)),
		molecule:     lines[len(lines)-1],
	}
	for i, s := range lines[:len(lines)-2] {
		parts := strings.Split(s, day19Separator)
		if len(parts) != 2 {
			return Day19Puzzle{}, fmt.Errorf("line %d: invalid replacement", i+1)
		}
		from := strings.TrimSpace(parts[0])
		into := strings.TrimSpace(parts[1])
		p.replacements[from] = append(p.replacements[from], into)
	}
	return p, nil
}

// Day19 solves day 19 for the selected part.
func Day19(p Day19Puzzle, part1 bool) uint {
	if part1 {
		return uint(len(day19Distinct(p.replacements, p.molecule)))
	}
	return day19ReduceSteps(p.molecule, day19Reducers(p.replacements))
}

func day19Distinct(replacements map[string][]string, molecule string) map[string]bool {
	m := make(map[string]bool)
	for from, intoList := range replacements {
		for _, into := range intoList {
			for c := strings.Count(molecule, from); c > 0; c-- {
				r := replaceNth(molecule, from, into, c)
				m[r] = true
			}
		}
	}
	return m
}

func day19Reducers(replacements map[string][]string) []day19Reducer {
	rs := make([]day19Reducer, 0, len(replacements))
	for from, intoList := range replacements {
		for _, into := range intoList {
			rs = append(rs, day19Reducer{from: into, into: from})
		}
	}
	slices.SortFunc(rs, func(a, b day19Reducer) int {
		if len(a.from) > len(b.from) {
			return -1
		}
		if len(a.from) < len(b.from) {
			return 1
		}
		return 0
	})
	return rs
}

func day19ReduceSteps(molecule string, rs []day19Reducer) uint {
	if steps := day19ReduceGreedy(molecule, rs); steps > 0 {
		return steps
	}
	return day19ReduceSearch(molecule, rs)
}

func day19ReduceGreedy(molecule string, rs []day19Reducer) uint {
	base := make([]day19Reducer, len(rs))
	copy(base, rs)
	rng := rand.New(rand.NewSource(1))
	for range 512 {
		order := make([]day19Reducer, len(base))
		copy(order, base)
		rng.Shuffle(len(order), func(i, j int) {
			order[i], order[j] = order[j], order[i]
		})
		current := molecule
		steps := uint(0)
		for current != "e" {
			progress := false
			for _, r := range order {
				i := strings.Index(current, r.from)
				if i < 0 {
					continue
				}
				current = current[:i] + r.into + current[i+len(r.from):]
				steps++
				progress = true
				break
			}
			if !progress {
				break
			}
		}
		if current == "e" {
			return steps
		}
	}
	return 0
}

func day19ReduceSearch(molecule string, rs []day19Reducer) uint {
	prospects := map[string]struct{}{molecule: {}}
	for step := uint(1); ; step++ {
		reduced := make(map[string]struct{})
		for p := range prospects {
			for _, r := range rs {
				for c := strings.Count(p, r.from); c > 0; c-- {
					s := replaceNth(p, r.from, r.into, c)
					reduced[s] = struct{}{}
				}
			}
		}
		if _, ok := reduced["e"]; ok {
			return step
		}
		if len(reduced) == 0 {
			return 0
		}

		ss := make([]string, 0, len(reduced))
		for k := range reduced {
			ss = append(ss, k)
		}
		slices.SortFunc(ss, func(a, b string) int {
			if len(a) < len(b) {
				return -1
			}
			if len(a) > len(b) {
				return 1
			}
			return 0
		})

		next := make(map[string]struct{})
		const heuristic = 100
		limit := min(len(ss), heuristic)
		for _, s := range ss[:limit] {
			next[s] = struct{}{}
		}
		prospects = next
	}
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
