package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

// TSP is the number of table spoon units for a combination of ingredients.
const (
	CALORIES = 500
	TSP      = 100
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

func (a Ingredient) calories() int {
	return a.Properties["calories"]
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

// Tsp is the cooking unit one teaspoon.
type Tsp uint

// Serving is one single ingredient of a cookie recipe.
type Serving struct {
	Ingredient
	tsp Tsp
}

func (a Serving) calories() uint {
	c := a.Ingredient.calories() * int(a.tsp)
	if c < 0 {
		return 0
	}
	return uint(c)
}

// Cookie consists of ingredients in full-teaspoon units.
type Cookie []Serving

func (a Cookie) tsps() []Tsp {
	var tsps []Tsp
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

func (a Cookie) calories() uint {
	var c uint
	for _, i := range a {
		c += i.calories()
	}
	return c
}

// Day15Part1 returns fittest cookie for all combinations of ingredients.
// No hardcoded number of ingredients, otherwise five embedded loops will do the
// combination trick.
func Day15Part1(is []Ingredient) Cookie {
	return day15(is, func(a Cookie) bool {
		// do not filter anything
		return false
	})
}

// Day15Part2 returns fittest cookie for all combinations of ingredients that
// does not exceed 500 calories. No hardcoded number of ingredients, otherwise
// five embedded loops will do the combination trick.
func Day15Part2(is []Ingredient) Cookie {
	return day15(is, func(a Cookie) bool {
		// this is a filter function, so we ignore any non-500
		return a.calories() != CALORIES
	})
}

func day15(is []Ingredient, cookieFilter func(Cookie) bool) Cookie {
	// start combination generator
	ch := make(chan ([]int))
	go KCompositions(TSP, len(is), ch)

	var combinations [][]int
	for digits := range ch {
		combinations = append(combinations, digits)
	}
	var champ Cookie
	var highscore uint
	for _, digits := range combinations {
		var c Cookie
		for i, digit := range digits {
			c = append(c, Serving{is[i], Tsp(digit)})
		}
		sc := c.score()
		if sc > highscore && !cookieFilter(c) {
			highscore = sc
			champ = c
		}
	}
	return champ
}
