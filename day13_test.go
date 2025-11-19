package adventofcode2015

import "testing"

func TestDay13Example(t *testing.T) {
	const want = 330
	got, err := Day13Part1(exampleFilename(13))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want = 709
	got, err := Day13Part1(filename(13))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part2(t *testing.T) {
	const want = 668
	got, err := Day13Part2(filename(13))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	fname := filename(13)
	b.ResetTimer()
	for range b.N {
		_, _ = Day13Part1(fname)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	fname := filename(13)
	b.ResetTimer()
	for range b.N {
		_, _ = Day13Part2(fname)
	}
}
