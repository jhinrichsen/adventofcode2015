package adventofcode2015

import (
	"testing"
)

const day04Input = "bgvyzdsv"

func TestDay04Example1(t *testing.T) {
	const want = 609_043
	got := Day04Part1("abcdef")
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestDay04Example2(t *testing.T) {
	const want = 1_048_970
	got := Day04Part1("pqrstuv")
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestDay04Part1(t *testing.T) {
	const want = 254_575
	got := Day04Part1(day04Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 1_038_736
	got := Day04Part2(day04Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	const input = "bgvyzdsv"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day04Part1(input)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	const input = "bgvyzdsv"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day04Part2(input)
	}
}
