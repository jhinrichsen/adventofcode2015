package adventofcode2015

import (
	"fmt"
	"testing"
)

var day20Samples = []struct {
	in  uint
	out uint
}{
	{1, 10},
	{2, 30},
	{3, 40},
	{4, 70},
	{5, 60},
	{6, 120},
	{7, 80},
	{8, 150},
	{9, 130},
}

func TestDay20Examples(t *testing.T) {
	for _, tt := range day20Samples {
		id := fmt.Sprintf("%d", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := presents(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 776160
	got := Day20Part1()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Highest ranking algorithm.
func TestDay20Champ(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	const want = 776160
	got := day20Part1Champ(36_000_000)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Champ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day20Part1Champ(36_000_000)
	}
}

func BenchmarkDay20MyChamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day20MyChamp(36_000_000)
	}
}

func TestDay20Part2(t *testing.T) {
	const want = 786240
	got := Day20Part2()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
