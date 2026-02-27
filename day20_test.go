package adventofcode2015

import "testing"

func BenchmarkDay20Part1(b *testing.B) {
	benchWithParser(b, 20, true, NewDay20, Day20)
}

func BenchmarkDay20Part2(b *testing.B) {
	benchWithParser(b, 20, false, NewDay20, Day20)
}

func TestDay20Examples(t *testing.T) {
	tests := []struct {
		in   uint
		want uint
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
	for _, tt := range tests {
		if got := day20PresentsPart1(tt.in); got != tt.want {
			t.Fatalf("house %d: want %d but got %d", tt.in, tt.want, got)
		}
	}
}

func TestDay20Part1(t *testing.T) {
	testWithParser(t, 20, filename, true, NewDay20, Day20, uint(776_160))
}

func TestDay20Part2(t *testing.T) {
	testWithParser(t, 20, filename, false, NewDay20, Day20, uint(786_240))
}
