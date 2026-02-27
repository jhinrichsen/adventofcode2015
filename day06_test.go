package adventofcode2015

import "testing"

func BenchmarkDay06Part1(b *testing.B) {
	benchWithParser(b, 6, true, NewDay06, Day06)
}

func BenchmarkDay06Part2(b *testing.B) {
	benchWithParser(b, 6, false, NewDay06, Day06)
}

func TestDay06Part1(t *testing.T) {
	testWithParser(t, 6, filename, true, NewDay06, Day06, uint(400_410))
}

func TestDay06Part2(t *testing.T) {
	testWithParser(t, 6, filename, false, NewDay06, Day06, uint(15_343_601))
}
