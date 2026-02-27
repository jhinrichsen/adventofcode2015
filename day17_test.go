package adventofcode2015

import "testing"

func BenchmarkDay17Part1(b *testing.B) {
	benchWithParser(b, 17, true, NewDay17, Day17)
}

func BenchmarkDay17Part2(b *testing.B) {
	benchWithParser(b, 17, false, NewDay17, Day17)
}

func TestDay17Part1Example(t *testing.T) {
	example := []uint{20, 15, 10, 5, 5}
	if got := day17Count(25, example, true); got != 4 {
		t.Fatalf("want %d but got %d", 4, got)
	}
}

func TestDay17Part1(t *testing.T) {
	testWithParser(t, 17, filename, true, NewDay17, Day17, uint(1_304))
}

func TestDay17Part2Example(t *testing.T) {
	example := []uint{20, 15, 10, 5, 5}
	if got := day17Count(25, example, false); got != 3 {
		t.Fatalf("want %d but got %d", 3, got)
	}
}

func TestDay17Part2(t *testing.T) {
	testWithParser(t, 17, filename, false, NewDay17, Day17, uint(18))
}
