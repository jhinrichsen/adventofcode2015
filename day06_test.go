package adventofcode2015

import "testing"

func BenchmarkDay06Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(6))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Day06Part1(lines)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(6))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Day06Part2(lines)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestDay06Part1(t *testing.T) {
	const want = 400_410
	lines := linesFromFilename(t, filename(6))
	got, err := Day06Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part2(t *testing.T) {
	const want = 15_343_601
	lines := linesFromFilename(t, filename(6))
	got, err := Day06Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
