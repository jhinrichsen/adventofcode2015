package adventofcode2015

import (
	"bufio"
	"os"
	"testing"
)

func TestHexEscape(t *testing.T) {
	// "\xa8br\x8bjr\""
	buf := []byte{
		'"',
		'\\', 'x', 'a', '8',
		'b', 'r',
		'\\', 'x', '8', 'b',
		'j', 'r', '\\', '"', '"'}
	wantLen := 16
	gotLen := len(buf)
	if wantLen != gotLen {
		t.Fatalf("want len %d but got %d", wantLen, gotLen)
	}

	want := 7
	got := Day8Part1(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
func TestDay8Part1Example(t *testing.T) {
	test(t, "testdata/day8_example.txt", 12)
}

func TestDay8Part1(t *testing.T) {
	test(t, "testdata/day8.txt", 1371)
}

func test(t *testing.T, filename string, want int) {
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	// m >= n
	mm := 0
	nn := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Bytes()
		mm += len(line)
		nn += Day8Part1(line)
	}
	got := mm - nn
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
