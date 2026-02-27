package adventofcode2015

import "testing"

const day04Input = "bgvyzdsv"

func TestDay04Part1Example1(t *testing.T) {
	const want = 609_043
	got, err := Day04([]byte("abcdef"), true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestDay04Part1Example2(t *testing.T) {
	const want = 1_048_970
	got, err := Day04([]byte("pqrstuv"), true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	for b.Loop() {
		_, _ = Day04([]byte(day04Input), true)
	}
}

func TestDay04Part1(t *testing.T) {
	const want = 254_575
	got, err := Day04([]byte(day04Input), true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	for b.Loop() {
		_, _ = Day04([]byte(day04Input), false)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 1_038_736
	got, err := Day04([]byte(day04Input), false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
