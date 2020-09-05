package adventofcode2015

import (
	"testing"
)

func TestDay9Part1Example(t *testing.T) {
	const want = 605
	got, _, err := Day9(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9Part1(t *testing.T) {
	const want = 117
	got, _, err := Day9(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9Part2Example(t *testing.T) {
	const want = 982
	_, got, err := Day9(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay9Part2(t *testing.T) {
	const want = 909
	_, got, err := Day9(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
