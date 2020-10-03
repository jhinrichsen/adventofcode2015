package adventofcode2015

import (
	"fmt"
	"testing"
)

//    |    1         2         3         4         5         6
// ---+---------+---------+---------+---------+---------+---------+
//  1 | 20151125  18749137  17289845  30943339  10071777  33511524
//  2 | 31916031  21629792  16929656   7726640  15514188   4041754
//  3 | 16080970   8057251   1601130   7981243  11661866  16474243
//  4 | 24592653  32451966  21345942   9380097  10600672  31527494
//  5 |    77061  17552253  28094349   6899651   9250759  31663883
//  6 | 33071741   6796745  25397450  24659492   1534922  27995004

var day25ExampleCodes = []uint{
	0,
	20151125, // n = 1
	31916031, // n = 2
	18749137, // n = 3
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

var day25Examples = []struct {
	x, y, code uint
}{
	{1, 1, 20151125},
	{2, 3, 8057251},
	{2, 5, 17552253},
	{3, 3, 1601130},
}

func TestDay25Sequence(t *testing.T) {
	d := newDay25()
	for l := uint(len(day25ExampleCodes)); d.n < l; d.next() {
		want := day25ExampleCodes[d.n]
		got := d.code
		if want != got {
			t.Fatalf("n=%d: want %d but got %d", d.n, want, got)
		}
	}
}

func TestDay25Examples(t *testing.T) {
	for _, tt := range day25Examples {
		id := fmt.Sprintf("(%d/%d)", tt.x, tt.y)
		t.Run(id, func(t *testing.T) {
			want := tt.code
			got := Day25Part1(tt.x, tt.y)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay25Part1(b *testing.B) {
	const (
		x    = 3075
		y    = 2981
		want = 9132360
	)

	for i := 0; i < b.N; i++ {
		got := Day25Part1(x, y)
		if want != got {
			b.Fatalf("want %d but got %d", want, got)
		}
	}
}
