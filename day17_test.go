package adventofcode2015

import (
	"strconv"
	"testing"
)

func TestDay17Example(t *testing.T) {
	const want = 4
	const storage = 25
	capacities := []uint{
		20, 15, 10, 5, 5,
	}
	got := Day17(storage, capacities)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part1(t *testing.T) {
	const want = 1304
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	var ns []uint
	for _, line := range lines {
		n, err := strconv.ParseUint(line, 10, 32)
		if err != nil {
			t.Fatal(err)
		}
		ns = append(ns, uint(n))
	}
	got := Day17(150, ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
