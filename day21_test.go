package adventofcode2015

import (
	"testing"
)

var boss = player{
	hitPoints: 100,
	damage:    8,
	armor:     2,
}
var nakedMe = player{100, 0, 0}

func TestDay21ExamplePart1(t *testing.T) {
	const want = 0
	players := [...]player{
		{8, 5, 5},  // Player #1
		{12, 7, 2}, // Player #2
	}
	Day21Part1(&players)
	if players[0].hitPoints != 2 {
		t.Fatalf("player #1: want 2 hit points but got %d", players[0].hitPoints)
	}
	if players[1].hitPoints != 0 {
		t.Fatalf("player #2: want 0 hit points but got %d", players[1].hitPoints)
	}
}

func TestDay21Part1Naked(t *testing.T) {
	const want = false // we will not survive without stuff
	players := [...]player{
		nakedMe, // Player #1
		boss,    // Player #2
	}
	Day21Part1(&players)
	got := players[0].alive()
	if want != got {
		t.Fatalf("want %t but got %t", want, got)
	}
}

func TestDay21Combinations(t *testing.T) {
	// cheapest equipment: Dagger, no armor, no rings
	const want = 8
	cs := combinations()
	got := cs[0].cost
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part1(t *testing.T) {
	const want = 91
	var got uint
	for _, c := range combinations() {
		p := nakedMe
		p.armor = c.armor
		p.damage = c.damage
		players := [...]player{
			p,
			boss,
		}
		Day21Part1(&players)
		if players[0].alive() {
			got = c.cost
			break
		}
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
