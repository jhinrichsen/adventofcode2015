package adventofcode2015

import (
	"testing"
)

func TestNewSue(t *testing.T) {
	s := "Sue 475: trees: 2, cars: 7, akitas: 8"
	sue, err := newSue(s)
	if err != nil {
		t.Fatal(err)
	}
	if sue.Number != 475 {
		t.Fatalf("want 475 but got %d", sue.Number)
	}
	if len(sue.Properties) != 3 {
		t.Fatalf("want 3 properties but got %d", len(sue.Properties))
	}
	if sue.Properties["trees"] != 2 {
		t.Fatalf("want 2 trees but got %d", sue.Properties["trees"])
	}
	if sue.Properties["cars"] != 7 {
		t.Fatalf("want 7 trees but got %d", sue.Properties["cars"])
	}
	if sue.Properties["akitas"] != 8 {
		t.Fatalf("want 8 trees but got %d", sue.Properties["akitas"])
	}
}

func probe() map[string]uint {
	m := make(map[string]uint)
	m["children"] = 3
	m["cats"] = 7
	m["samoyeds"] = 2
	m["pomeranians"] = 3
	m["akitas"] = 0
	m["vizslas"] = 0
	m["goldfish"] = 5
	m["trees"] = 3
	m["cars"] = 2
	m["perfumes"] = 1
	return m
}

func TestDay16Part1(t *testing.T) {
	const want = 373
	m := probe()
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	var got uint
	for _, line := range lines {
		sue, err := newSue(line)
		if err != nil {
			t.Fatal(err)
		}
		if match(m, sue.Properties) {
			got = sue.Number
			break
		}
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part2(t *testing.T) {
	const want = 260
	m := probe()
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	var got uint
	for _, line := range lines {
		sue, err := newSue(line)
		if err != nil {
			t.Fatal(err)
		}
		if matchWorkaroundForBrokenTurboEncapulator(m, sue.Properties) {
			got = sue.Number
			break
		}
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
