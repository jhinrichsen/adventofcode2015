package adventofcode2015

import "testing"

func BenchmarkDay13Part1(b *testing.B) {
	benchWithParser(b, 13, true, NewDay13, Day13)
}

func BenchmarkDay13Part2(b *testing.B) {
	benchWithParser(b, 13, false, NewDay13, Day13)
}

func TestDay13Part1Example(t *testing.T) {
	testWithParser(t, 13, exampleFilename, true, NewDay13, Day13, uint(330))
}

func TestDay13Part1(t *testing.T) {
	testWithParser(t, 13, filename, true, NewDay13, Day13, uint(709))
}

func TestDay13Part2(t *testing.T) {
	testWithParser(t, 13, filename, false, NewDay13, Day13, uint(668))
}

