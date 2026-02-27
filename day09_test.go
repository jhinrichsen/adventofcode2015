package adventofcode2015

import "testing"

func TestDay9Part1Example(t *testing.T) {
	const want = 605
	lines := linesFromFilename(t, exampleFilename(9))
	got, _, err := Day9(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9Part1(t *testing.T) {
	const want = 117
	lines := linesFromFilename(t, filename(9))
	got, _, err := Day9(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9Part2Example(t *testing.T) {
	const want = 982
	lines := linesFromFilename(t, exampleFilename(9))
	_, got, err := Day9(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9Part2(t *testing.T) {
	const want = 909
	lines := linesFromFilename(t, filename(9))
	_, got, err := Day9(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
