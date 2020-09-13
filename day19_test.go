package adventofcode2015

import (
	"reflect"
	"sort"
	"testing"
)

func hhoPlant() plant {
	p, _ := newPlant([]string{
		"H => HO",
		"H => OH",
		"O => HH"})
	return p
}

func TestDay19AddReplacement(t *testing.T) {
	p := hhoPlant()

	if len(p.replacements) != 2 {
		t.Fatalf("want len(2) but got len(%d)", len(p.replacements))
	}
	if len(p.replacements["O"]) != 1 {
		t.Fatalf("want 1 replacement for O but got %d",
			len(p.replacements["O"]))
	}
	if len(p.replacements["H"]) != 2 {
		t.Fatalf("want 2 replacements for H but got %d",
			len(p.replacements["H"]))
	}
}

func TestDay19Hoh(t *testing.T) {
	const want = 4
	p := hhoPlant()
	p.molecule = "HOH"
	got := Day19Part1(p)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Hohoho(t *testing.T) {
	const want = 7
	p := hhoPlant()
	p.molecule = "HOHOHO"
	got := Day19Part1(p)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part1(t *testing.T) {
	const want = 576
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}
	p, err := newPlant(lines[:len(lines)-2])
	if err != nil {
		t.Fatal(err)
	}
	p.molecule = lines[len(lines)-1]
	got := Day19Part1(p)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func part2ExamplePlant() plant {
	rs := []string{
		"e => H",
		"e => O",
		"H => HO",
		"H => OH",
		"O => HH",
	}
	p, _ := newPlant(rs)
	return p
}

func TestDay19Example1Part2(t *testing.T) {
	const want = 3
	p := part2ExamplePlant()
	p.molecule = "HOH"
	got := Day19Part2(p.molecule, p.reducers())
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Example2Part2(t *testing.T) {
	const want = 6
	p := part2ExamplePlant()
	p.molecule = "HOHOHO"
	got := Day19Part2(p.molecule, p.reducers())
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2(t *testing.T) {
	const want = 207
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}
	p, err := newPlant(lines[:len(lines)-2])
	if err != nil {
		t.Fatal(err)
	}
	p.molecule = lines[len(lines)-1]
	got := Day19Part2(p.molecule, p.reducers())
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestSort(t *testing.T) {
	want := []string{
		"e", "bb", "aaa", "ccc", "aaaaa", "ddddddddd",
	}
	got := []string{
		"aaa", "aaaaa", "bb", "ccc", "ddddddddd", "e",
	}
	// make sure _not_ to use sort.Strings(ByLen(got))
	sort.Sort(ByLen(got))
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}
