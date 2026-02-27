package adventofcode2015

import (
	"fmt"
	"testing"
)

func BenchmarkDay25Part1(b *testing.B) {
	benchWithParser(b, 25, true, NewDay25, Day25)
}

func TestDay25Sequence(t *testing.T) {
	exampleCodes := []uint{
		0,
		20151125,
		31916031,
		18749137,
		16080970,
		21629792,
		17289845,
		24592653,
		8057251,
		16929656,
		30943339,
		77061,
		32451966,
	}
	d := newDay25State()
	for l := uint(len(exampleCodes)); d.n < l; d.next() {
		if got := d.code; got != exampleCodes[d.n] {
			t.Fatalf("n=%d: want %d but got %d", d.n, exampleCodes[d.n], got)
		}
	}
}

func TestDay25Examples(t *testing.T) {
	tests := []struct {
		x, y, code uint
	}{
		{1, 1, 20151125},
		{2, 3, 8057251},
		{2, 5, 17552253},
		{3, 3, 1601130},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%d/%d)", tt.x, tt.y), func(t *testing.T) {
			if got := day25CodeAt(tt.x, tt.y); got != tt.code {
				t.Fatalf("want %d but got %d", tt.code, got)
			}
		})
	}
}

func TestDay25Part1(t *testing.T) {
	testWithParser(t, 25, filename, true, NewDay25, Day25, uint(9_132_360))
}
