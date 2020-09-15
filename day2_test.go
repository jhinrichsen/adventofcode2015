package adventofcode2015

import (
	"fmt"
	"testing"
)

var day2ExamplesPart1 = []struct {
	in  string
	out uint
}{
	{"2x3x4", 58},
	{"1x1x10", 43},
}

func TestDay2ExamplesPart1(t *testing.T) {
	for _, tt := range day2ExamplesPart1 {
		id := fmt.Sprintf("%+v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day2Part1([]string{tt.in})
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay2Part1(t *testing.T) {
	const want = 1_598_415
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day2Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

var day2ExamplesPart2 = []struct {
	in  string
	out uint
}{
	{"2x3x4", 34},
	{"1x1x10", 14},
}

func TestDay2ExamplesPart2(t *testing.T) {
	for _, tt := range day2ExamplesPart2 {
		id := fmt.Sprintf("%+v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day2Part2([]string{tt.in})
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
	const want = 3_812_909
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day2Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
