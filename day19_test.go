package adventofcode2015

import "testing"

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
	got := p.distinct()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Hohoho(t *testing.T) {
	const want = 7
	p := hhoPlant()
	p.molecule = "HOHOHO"
	got := p.distinct()
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
	got := p.distinct()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
