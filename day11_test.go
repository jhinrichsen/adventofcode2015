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

func inc(s string) string {
	bs := []byte(s)
	idx := len(bs) - 1
_inc:
	if bs[idx] == 'z' {
		bs[idx] = 'a'
		idx--
		goto _inc
	} else {
		bs[idx]++
	}
	return string(bs)
}
