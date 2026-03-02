package adventofcode2015

import "testing"

func BenchmarkDay11Part1(b *testing.B) {
	benchWithParser(b, 11, true, NewDay11, Day11)
}

func BenchmarkDay11Part2(b *testing.B) {
	benchWithParser(b, 11, false, NewDay11, Day11)
}

func TestDay11Inc(t *testing.T) {
	bs := []byte("xx")
	incBytes(bs)
	if got := string(bs); got != "xy" {
		t.Fatalf("want %q but got %q", "xy", got)
	}
	incBytes(bs)
	if got := string(bs); got != "xz" {
		t.Fatalf("want %q but got %q", "xz", got)
	}
	incBytes(bs)
	if got := string(bs); got != "ya" {
		t.Fatalf("want %q but got %q", "ya", got)
	}
	incBytes(bs)
	if got := string(bs); got != "yb" {
		t.Fatalf("want %q but got %q", "yb", got)
	}
}

func TestDay11hijklmn(t *testing.T) {
	bs := []byte("hijklmn")
	if !req1Bytes(bs) {
		s := string(bs)
		t.Fatalf("%q does not meet requirement #1", s)
	}
	if req2Bytes(bs) {
		s := string(bs)
		t.Fatalf("%q does meet requirement #2", s)
	}
}

func TestDay11abbceffg(t *testing.T) {
	bs := []byte("abbceffg")
	if req1Bytes(bs) {
		s := string(bs)
		t.Fatalf("%q does meet requirement #1", s)
	}
	if !req3Bytes(bs) {
		s := string(bs)
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
	testWithParser(t, 11, filename, true, NewDay11, Day11, "hepxxyzz")
}

func TestDay11Part2(t *testing.T) {
	testWithParser(t, 11, filename, false, NewDay11, Day11, "heqaabcc")
}
