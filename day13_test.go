package adventofcode2015

import "testing"

func TestDay13Example(t *testing.T) {
	const want = 330
	got, err := Day13(exampleFilename(13))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want = 709
	got, err := Day13(filename(13))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
