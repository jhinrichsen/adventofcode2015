package adventofcode2015

import "testing"

func TestDay13Example(t *testing.T) {
	const want = 330
	lines := linesFromFilename(t, exampleFilename(13))
	got, err := Day13Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want = 709
	lines := linesFromFilename(t, filename(13))
	got, err := Day13Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part2(t *testing.T) {
	const want = 668
	lines := linesFromFilename(t, filename(13))
	got, err := Day13Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
