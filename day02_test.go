package adventofcode2015

import (
	"fmt"
	"testing"
)

func TestDay02ExamplesPart1(t *testing.T) {
	day2ExamplesPart1 := []struct {
		in  string
		out uint
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}
	for _, tt := range day2ExamplesPart1 {
		id := fmt.Sprintf("%+v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			puzzle, err := NewDay02([]string{tt.in})
			if err != nil {
				t.Fatal(err)
			}
			got := Day02(puzzle, true)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	benchWithParser(b, 2, true, NewDay02, Day02)
}

func TestDay02Part1(t *testing.T) {
	testWithParser(t, 2, filename, true, NewDay02, Day02, uint(1_598_415))
}

func TestDay02ExamplesPart2(t *testing.T) {
	day2ExamplesPart2 := []struct {
		in  string
		out uint
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}
	for _, tt := range day2ExamplesPart2 {
		id := fmt.Sprintf("%+v", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			puzzle, err := NewDay02([]string{tt.in})
			if err != nil {
				t.Fatal(err)
			}
			got := Day02(puzzle, false)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	benchWithParser(b, 2, false, NewDay02, Day02)
}

func TestDay02Part2(t *testing.T) {
	testWithParser(t, 2, filename, false, NewDay02, Day02, uint(3_812_909))
}
