package adventofcode2015

import "testing"

func BenchmarkDay10Part1(b *testing.B) {
	benchWithParser(b, 10, true, NewDay10, Day10)
}

func BenchmarkDay10Part2(b *testing.B) {
	benchWithParser(b, 10, false, NewDay10, Day10)
}

func TestDay10Samples(t *testing.T) {
	samples := []struct {
		in  string
		out string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			if got := lookAndSay(tt.in); got != tt.out {
				t.Fatalf("%q: want %q but got %q", tt.in, tt.out, got)
			}
		})
	}
}

func TestDay10Part1(t *testing.T) {
	testWithParser(t, 10, filename, true, NewDay10, Day10, uint(360_154))
}

func TestDay10Part2(t *testing.T) {
	testWithParser(t, 10, filename, false, NewDay10, Day10, uint(5_103_798))
}
