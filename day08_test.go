package adventofcode2015

import "testing"

func BenchmarkDay08Part1(b *testing.B) {
	benchLines(b, 8, true, Day08)
}

func BenchmarkDay08Part2(b *testing.B) {
	benchLines(b, 8, false, Day08)
}

func TestDay08Part1Example(t *testing.T) {
	testLines(t, 8, exampleFilename, true, Day08, uint(12))
}

func TestDay08Part1(t *testing.T) {
	testLines(t, 8, filename, true, Day08, uint(1_371))
}

func TestDay08Part2Example(t *testing.T) {
	testLines(t, 8, exampleFilename, false, Day08, uint(19))
}

func TestDay08Part2(t *testing.T) {
	testLines(t, 8, filename, false, Day08, uint(2_117))
}
