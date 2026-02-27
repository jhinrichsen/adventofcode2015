package adventofcode2015

import (
	"strconv"
	"testing"
)

const storage = 150

func TestDay17Example(t *testing.T) {
	exampleCapacities := []uint{
		20, 15, 10, 5, 5,
	}
	const want = 4
	const storage = 25
	got := Day17Part1(storage, exampleCapacities)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func numbers(tb testing.TB, filename string) ([]uint, error) {
	var ns []uint
	lines := linesFromFilename(tb, filename)
	for _, line := range lines {
		n, err := strconv.ParseUint(line, 10, 32)
		if err != nil {
			return ns, err
		}
		ns = append(ns, uint(n))
	}
	return ns, nil
}

func TestDay17Part1(t *testing.T) {
	const want = 1304
	ns, err := numbers(t, filename(17))
	if err != nil {
		t.Fatal(err)
	}
	got := Day17Part1(storage, ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part2Example(t *testing.T) {
	exampleCapacities := []uint{
		20, 15, 10, 5, 5,
	}
	const want = 3
	got := Day17Part2(25, exampleCapacities)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part2(t *testing.T) {
	const want = 18
	ns, err := numbers(t, filename(17))
	if err != nil {
		t.Fatal(err)
	}
	got := Day17Part2(storage, ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
