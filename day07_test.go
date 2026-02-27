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
		got, err := puzzle.signal(wire, nil)
		if err != nil {
			t.Fatal(err)
		}
		if got != signal {
			t.Fatalf("%s: want %d but got %d", wire, signal, got)
		}
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	benchLinesErr(b, 7, true, func(lines []string, part1 bool) (uint16, error) {
		if part1 {
			return Day07Part1(lines)
		}
		return Day07Part2(lines)
	})
}

func BenchmarkDay07Part2(b *testing.B) {
	benchLinesErr(b, 7, false, func(lines []string, part1 bool) (uint16, error) {
		if part1 {
			return Day07Part1(lines)
		}
		return Day07Part2(lines)
	})
}

func TestDay07Part1(t *testing.T) {
	const want = 16076
	lines := linesFromFilename(t, filename(7))
	got, err := Day07Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 2797
	lines := linesFromFilename(t, filename(7))
	got, err := Day07Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
