package adventofcode2015

import (
	"testing"
)

func TestDay21Part1Naked(t *testing.T) {
	const want = false // we will not survive without stuff
	players := [...]day21Player{
		nakedMe,   // Player #1
		day21Boss, // Player #2
	}
	duel(&players)
	got := players[0].alive()
	if want != got {
		t.Fatalf("want %t but got %t", want, got)
	}
}

func TestDay21Part1(t *testing.T) {
	const want = 91
	got := Day21Part1()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part2(t *testing.T) {
	const want = 158
	got := Day21Part2()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
