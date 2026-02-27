package adventofcode2015

import "testing"

func BenchmarkDay21Part1(b *testing.B) {
	benchWithParser(b, 21, true, NewDay21, Day21)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchWithParser(b, 21, false, NewDay21, Day21)
}

func TestDay21NakedLoses(t *testing.T) {
	boss := day21Player{hitPoints: 100, damage: 8, armor: 2}
	me := day21Player{hitPoints: 100, damage: 0, armor: 0}
	if day21Win(me, boss) {
		t.Fatalf("expected naked player to lose")
	}
}

func TestDay21Part1(t *testing.T) {
	testWithParser(t, 21, filename, true, NewDay21, Day21, uint(91))
}

func TestDay21Part2(t *testing.T) {
	testWithParser(t, 21, filename, false, NewDay21, Day21, uint(158))
}
