package adventofcode2015

import "testing"

func BenchmarkDay15Part1(b *testing.B) {
	benchWithParser(b, 15, true, NewDay15, Day15)
}

func BenchmarkDay15Part2(b *testing.B) {
	benchWithParser(b, 15, false, NewDay15, Day15)
}

func TestDay15Part1Example(t *testing.T) {
	testWithParser(t, 15, exampleFilename, true, NewDay15, Day15, uint(62_842_880))
}

func TestDay15Part2Example(t *testing.T) {
	testWithParser(t, 15, exampleFilename, false, NewDay15, Day15, uint(57_600_000))
}

func TestDay15Part1(t *testing.T) {
	testWithParser(t, 15, filename, true, NewDay15, Day15, uint(13_882_464))
}

func TestDay15Part2(t *testing.T) {
	testWithParser(t, 15, filename, false, NewDay15, Day15, uint(11_171_160))
}

