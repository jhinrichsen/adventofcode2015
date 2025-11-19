package adventofcode2015

import "testing"

func TestDay23Example(t *testing.T) {
	const want = 2
	instructions := linesFromFilename(t, exampleFilename(23))
	got, _ := Day23Part1(instructions)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	const want = 255
	is := linesFromFilename(t, filename(23))
	_, got := Day23Part1(is)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	is := linesFromFilename(b, filename(23))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day23Part1(is)
	}
}

func TestDay23Part2(t *testing.T) {
	const want = 334
	is := linesFromFilename(t, filename(23))
	_, got := Day23Part2(is, 1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	is := linesFromFilename(b, filename(23))
	b.ResetTimer()
	for range b.N {
		_, _ = Day23Part2(is, 1)
	}
}
