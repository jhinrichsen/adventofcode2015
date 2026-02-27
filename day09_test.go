package adventofcode2015

import "testing"

func BenchmarkDay09Part1(b *testing.B) {
	benchWithParser(b, 9, true, NewDay09, Day09)
}

func BenchmarkDay09Part2(b *testing.B) {
	benchWithParser(b, 9, false, NewDay09, Day09)
}

func TestDay09Part1Example(t *testing.T) {
	testWithParser(t, 9, exampleFilename, true, NewDay09, Day09, uint(605))
}

func TestDay09Part1(t *testing.T) {
	testWithParser(t, 9, filename, true, NewDay09, Day09, uint(117))
}

func TestDay09Part2Example(t *testing.T) {
	testWithParser(t, 9, exampleFilename, false, NewDay09, Day09, uint(982))
}

func TestDay09Part2(t *testing.T) {
	testWithParser(t, 9, filename, false, NewDay09, Day09, uint(909))
}
