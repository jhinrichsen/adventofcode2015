package adventofcode2015

import (
	"testing"
)

const day4Input = "bgvyzdsv"

func TestDay4Example1(t *testing.T) {
	const want = 609_043
	got := Day4Part1("abcdef")
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestDay4Example2(t *testing.T) {
	const want = 1_048_970
	got := Day4Part1("pqrstuv")
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestDay4Part1(t *testing.T) {
	const want = 254_575
	got := Day4Part1(day4Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay4Part2(t *testing.T) {
	const want = 1_038_736
	got := Day4Part2(day4Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
