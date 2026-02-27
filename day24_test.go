package adventofcode2015

import "testing"

func TestDay24ExamplePart1(t *testing.T) {
	const want = 99
	lines := linesFromFilename(t, exampleFilename(24))
	ws, err := newWeights(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day24Part1(ws)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	if testing.Short() {
		t.Skip("billions of permutations, will eventually finish")
	}
	const want = 11266889531
	lines := linesFromFilename(t, filename(24))
	ws, err := newWeights(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day24Part1(ws)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24ExamplePart2(t *testing.T) {
	const want = 44
	lines := linesFromFilename(t, exampleFilename(24))
	ws, err := newWeights(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day24Part2(ws)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part2(t *testing.T) {
	if testing.Short() {
		t.Skip("billions of permutations, will eventually finish")
	}
	const want = 77387711
	lines := linesFromFilename(t, filename(24))
	ws, err := newWeights(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day24Part1(ws)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
