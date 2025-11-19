package adventofcode2015

import "testing"

func TestDay11Inc(t *testing.T) {
	s := "xx"
	s = inc(s)
	if s != "xy" {
		t.Fatalf("want %q but got %q", "xy", s)
	}
	s = inc(s)
	if s != "xz" {
		t.Fatalf("want %q but got %q", "xz", s)
	}
	s = inc(s)
	if s != "ya" {
		t.Fatalf("want %q but got %q", "ya", s)
	}
	s = inc(s)
	if s != "yb" {
		t.Fatalf("want %q but got %q", "yb", s)
	}
}

func TestDay11hijklmn(t *testing.T) {
	const s = "hijklmn"
	if !req1(s) {
		t.Fatalf("%q does not meet requirement #1", s)
	}
	if req2(s) {
		t.Fatalf("%q does meet requirement #2", s)
	}
}

func TestDay11abbceffg(t *testing.T) {
	const s = "abbceffg"
	if req1(s) {
		t.Fatalf("%q does meet requirement #1", s)
	}
	if !req3(s) {
		t.Fatalf("%q does not meet requirement #3", s)
	}

}

func TestDay11Nextabcdefgh(t *testing.T) {
	const want = "abcdffaa"
	got := next("abcdefgh")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay11Nextghijklmn(t *testing.T) {
	const want = "ghjaabcc"
	got := next("ghijklmn")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay11Part1(t *testing.T) {
	const want = "hepxxyzz"
	got := next("hepxcrrq")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay11Part2(t *testing.T) {
	const want = "heqaabcc"
	got := next(next("hepxcrrq"))
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		_ = next("hepxcrrq")
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		_ = next(next("hepxcrrq"))
	}
}
