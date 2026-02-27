package adventofcode2015

import "testing"

func BenchmarkDay23Part1(b *testing.B) {
	benchWithParser(b, 23, true, NewDay23, Day23)
}

func BenchmarkDay23Part2(b *testing.B) {
	benchWithParser(b, 23, false, NewDay23, Day23)
}

func TestDay23Example(t *testing.T) {
	puzzle, err := NewDay23(linesFromFilename(t, exampleFilename(23)))
	if err != nil {
		t.Fatal(err)
	}
	a, _ := day23Run(puzzle, 0, 0)
	if a != 2 {
		t.Fatalf("want %d but got %d", 2, a)
	}
}

func TestDay23Part1(t *testing.T) {
	testWithParser(t, 23, filename, true, NewDay23, Day23, uint(255))
}

func TestDay23Part2(t *testing.T) {
	testWithParser(t, 23, filename, false, NewDay23, Day23, uint(334))
}

