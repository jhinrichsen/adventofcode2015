package adventofcode2015

import "testing"

func TestDay24Part1Example(t *testing.T) {
	testWithParser(t, 24, exampleFilename, true, NewDay24, Day24, uint(99))
}

func TestDay24Part1(t *testing.T) {
	testWithParser(t, 24, filename, true, NewDay24, Day24, uint(11_266_889_531))
}

func TestDay24Part2Example(t *testing.T) {
	testWithParser(t, 24, exampleFilename, false, NewDay24, Day24, uint(44))
}

func TestDay24Part2(t *testing.T) {
	testWithParser(t, 24, filename, false, NewDay24, Day24, uint(77_387_711))
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParser(b, 24, true, NewDay24, Day24)
}

func BenchmarkDay24Part2(b *testing.B) {
	benchWithParser(b, 24, false, NewDay24, Day24)
}
