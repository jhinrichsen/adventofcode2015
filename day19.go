package adventofcode2015

import (
	"fmt"
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

func (a plant) distinct() uint {
	m := make(map[string]bool)
	for from, v := range a.replacements {
		for _, into := range v {
			for c := strings.Count(a.molecule, from); c > 0; c-- {
				r := ReplaceNth(a.molecule, from, into, c)
				m[r] = true
			}
		}
	}
	return uint(len(m))
}
