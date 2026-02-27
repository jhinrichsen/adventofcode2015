package adventofcode2015

import "testing"

func TestDay07Example(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(7))
	if err != nil {
		t.Fatal(err)
	}
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
		got, err := puzzle.Signal(wire)
		if err != nil {
			t.Fatal(err)
		}
		if got != signal {
			t.Fatalf("%s: want %d but got %d", wire, signal, got)
		}
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Day07Part1(lines)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Day07Part2(lines)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 16076
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		t.Fatal(err)
	}
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
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
