package adventofcode2015

import "testing"

func BenchmarkDay14Part1(b *testing.B) {
	benchWithParser(b, 14, true, NewDay14, Day14)
}

func BenchmarkDay14Part2(b *testing.B) {
	benchWithParser(b, 14, false, NewDay14, Day14)
}

func TestDay14Part1Example(t *testing.T) {
	puzzle, err := NewDay14(linesFromFilename(t, exampleFilename(14)))
	if err != nil {
		t.Fatal(err)
	}
	if got := day14DistanceWinner(puzzle, 1000); got != 1120 {
		t.Fatalf("want %d but got %d", 1120, got)
	}
}

func TestDay14Part2Example(t *testing.T) {
	puzzle, err := NewDay14(linesFromFilename(t, exampleFilename(14)))
	if err != nil {
		t.Fatal(err)
	}
	if got := day14ScoreWinner(puzzle, 1000); got != 689 {
		t.Fatalf("want %d but got %d", 689, got)
	}
}

func TestDay14Part1(t *testing.T) {
	testWithParser(t, 14, filename, true, NewDay14, Day14, uint(2_655))
}

func TestDay14Part2(t *testing.T) {
	testWithParser(t, 14, filename, false, NewDay14, Day14, uint(1_059))
}
