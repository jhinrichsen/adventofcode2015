package adventofcode2015

import "testing"

func TestDay23Example(t *testing.T) {
	const want = 2
	instructions, err := linesFromFilename(exampleFilename(23))
	if err != nil {
		t.Fatal(err)
	}
	got, _ := Day23Part1(instructions)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	const want = 255
	is, err := linesFromFilename(filename(23))
	if err != nil {
		t.Fatal(err)
	}
	_, got := Day23Part1(is)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	is, err := linesFromFilename(filename(23))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day23Part1(is)
	}
}

func TestDay23Part2(t *testing.T) {
	const want = 334
	is, err := linesFromFilename(filename(23))
	if err != nil {
		t.Fatal(err)
	}
	_, got := Day23Part2(is, 1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
