package adventofcode2015

import "testing"

func TestDay07Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(7))
	puzzle, err := NewDay07(lines)
	if err != nil {
		t.Fatal(err)
	}

	want := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}
	for wire, signal := range want {
		id, err := day07WireID(wire)
		if err != nil {
			t.Fatal(err)
		}
		got, err := puzzle.signal(id, nil)
		if err != nil {
			t.Fatal(err)
		}
		if got != signal {
			t.Fatalf("%s: want %d but got %d", wire, signal, got)
		}
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	benchWithParser(b, 7, true, NewDay07, Day07)
}

func BenchmarkDay07Part2(b *testing.B) {
	benchWithParser(b, 7, false, NewDay07, Day07)
}

func TestDay07Part1(t *testing.T) {
	testWithParser(t, 7, filename, true, NewDay07, Day07, uint(16_076))
}

func TestDay07Part2(t *testing.T) {
	testWithParser(t, 7, filename, false, NewDay07, Day07, uint(2_797))
}
