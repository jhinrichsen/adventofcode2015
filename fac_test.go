package adventofcode2015

import (
	"fmt"
	"testing"
)

func TestFac(t *testing.T) {
	facSamples := []struct {
		in, out uint
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}
	for _, tt := range facSamples {
		id := fmt.Sprintf("%d", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Fac(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}
