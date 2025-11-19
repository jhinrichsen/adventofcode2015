package adventofcode2015

import (
	"fmt"
	"testing"
)

var day02ExamplesPart1 = []struct {
	in  string
	out uint
}{
	{"2x3x4", 58},
	{"1x1x10", 43},
}

func TestDay02ExamplesPart1(t *testing.T) {
	for _, tt := range day02ExamplesPart1 {
		id := fmt.Sprintf("%+v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day02Part1([]string{tt.in})
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 1_598_415
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

var day02ExamplesPart2 = []struct {
	in  string
	out uint
}{
	{"2x3x4", 34},
	{"1x1x10", 14},
}

func TestDay02ExamplesPart2(t *testing.T) {
	for _, tt := range day02ExamplesPart2 {
		id := fmt.Sprintf("%+v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day02Part2([]string{tt.in})
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay02Part2(t *testing.T) {
	const want = 3_812_909
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
