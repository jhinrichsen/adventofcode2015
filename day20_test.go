package adventofcode2015

import (
	"fmt"
	"testing"
)

var day20Samples = []struct {
	houseno  uint
	presents uint
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
		id := fmt.Sprintf("%d", tt.houseno)
		t.Run(id, func(t *testing.T) {
			want := tt.presents
			got := presents(tt.houseno)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 776160
	got := Day20(Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
