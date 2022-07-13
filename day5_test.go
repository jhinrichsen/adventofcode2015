package adventofcode2015

import (
	"testing"
)

var examplesDay5 = []struct {
	in  string
	f   func(string) bool
	out uint
}{
	{"ugknbfddgicrmopn", nicePart1, 1}, // 1 == Nice
	{"aaa", nicePart1, 1},
	{"jchzalrnumimnmhp", nicePart1, 0},
	{"haegwjzuvuyypxyu", nicePart1, 0},
	{"dvszwmarrgswjxmb", nicePart1, 0},

	{"qjhvhtzxzqqjkmpb", nicePart2, 1},
	{"xxyxx", nicePart2, 1},
	{"uurcxstgmygtbstg", nicePart2, 0},
	{"ieodomkazucvgmuy", nicePart2, 0},
}

func TestDay5Example1(t *testing.T) {
	for _, tt := range examplesDay5 {
		t.Run(tt.in, func(t *testing.T) {
			want := tt.out
			got := day5([]string{tt.in}, tt.f)
			if want != got {
				t.Fatalf("want %v but got %v", want, got)
			}
		})
	}
}

func TestDay5Part1(t *testing.T) {
	const want = 238
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got := day5(lines, nicePart1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay5Part2(t *testing.T) {
	const want = 69
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got := day5(lines, nicePart2)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
