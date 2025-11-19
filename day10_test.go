package adventofcode2015

import (
	"testing"
)

const (
	inputDay10 = "1113122113"
)

var samples = []struct {
	in  string
	out string
}{
	{"1", "11"},
	{"11", "21"},
	{"21", "1211"},
	{"1211", "111221"},
	{"111221", "312211"},
}

func TestDay10Samples(t *testing.T) {
	for _, tt := range samples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := lookAndSay(tt.in)
			if want != got {
				t.Fatalf("%q: want %q but got %q", id, want, got)
			}
		})
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 360154
	s := inputDay10
	for i := 0; i < 40; i++ {
		s = lookAndSay(s)
	}
	got := len(s)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2(t *testing.T) {
	const want = 5103798
	s := inputDay10
	for i := 0; i < 50; i++ {
		s = lookAndSay(s)
	}
	got := len(s)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		s := inputDay10
		for range 40 {
			s = lookAndSay(s)
		}
		_ = len(s)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		s := inputDay10
		for range 50 {
			s = lookAndSay(s)
		}
		_ = len(s)
	}
}
