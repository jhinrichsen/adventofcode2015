package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

// Ingredient are what cookies consist of.
type Ingredient struct {
	Name       string
	Properties map[string]int
}

// properties without calories.
func (a Ingredient) values() []int {
	return []int{
		a.Properties["capacity"],
		a.Properties["durability"],
		a.Properties["flavor"],
		a.Properties["texture"],
		// a.Properties["calories"],
	}
}

// NewIngredient parses a line in form "Butterscotch: capacity -1, durability
// -2, flavor 6, texture 3, calories 8" into an Ingredient.
func NewIngredient(s string) (Ingredient, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return Ingredient{},
			fmt.Errorf("want %d parts but got %d", 2, len(parts))
	}
	props := strings.Split(parts[1], ",")
	if len(props) != 5 {
		return Ingredient{},
			fmt.Errorf("want %d parts but got %d", 5, len(props))
	}
	m := make(map[string]int)
	for _, prop := range props {
		ps := strings.Fields(prop)
		n, err := strconv.Atoi(ps[1])
		if err != nil {
			return Ingredient{},
				fmt.Errorf("not a number for property %q: %q",
					ps[0], ps[1])
		}
		m[ps[0]] = n
	}
	return Ingredient{parts[0], m}, nil
}

// tsp is the cooking unit one teaspoon.
type tsp uint

type Serving struct {
	Ingredient
	tsp tsp
}

// Cookie consists of ingredients in full-teaspoon units.
type Cookie []Serving

func (a Cookie) tsps() []tsp {
	var tsps []tsp
	for _, s := range a {
		tsps = append(tsps, s.tsp)
	}
	return tsps
}

// properties without calories.
func (a Cookie) properties() [][]int {
	ps := make([][]int, len(a))
	for y := range a {
		ps[y] = a[y].values()
	}
	return ps
}

func (a Cookie) score() uint {
	product := uint(1)
	tsps := a.tsps()        // list of tsps
	props := a.properties() // matrix of properties, y same order as tsps
	for x := range props[0] {
		propTotal := 0
		for y := range a {
			propTotal += int(tsps[y]) * props[y][x]
		}
		// If any properties had produced a negative total, it would
		// have instead become zero
		if propTotal < 0 {
			propTotal = 0
		}
		product *= uint(propTotal)
	}
	return product
}
