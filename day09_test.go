package adventofcode2015

import (
	"testing"
)

func TestDay09Part1Example(t *testing.T) {
	const want = 605
	got, _, err := Day09(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part1(t *testing.T) {
	const want = 117
	got, _, err := Day09(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part2Example(t *testing.T) {
	const want = 982
	_, got, err := Day09(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part2(t *testing.T) {
	const want = 909
	_, got, err := Day09(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day09(filename(9))
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day09(filename(9))
	}
}
