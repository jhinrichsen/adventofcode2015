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
	got := Day08Part1(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
func TestDay08Part1Example(t *testing.T) {
	test(t, exampleFilename(8), 12)
}

func TestDay08Part1(t *testing.T) {
	test(t, filename(8), 1371)
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
		nn += Day08Part1(line)
	}
	got := mm - nn
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

var part2Tests = []struct {
	in  []byte
	out int
}{
	{[]byte{'"', '"'}, 6},
	{[]byte{'"', 'a', 'b', 'c', '"'}, 9},
	{[]byte{'"', 'a', 'a', 'a', '\\', '"', 'a', 'a', 'a', '"'}, 16},
	{[]byte{'"', '\\', 'x', '2', '7', '"'}, 11},
}

func TestDay08Part2Examples(t *testing.T) {
	for _, tt := range part2Tests {
		t.Run(string(tt.in), func(t *testing.T) {
			want := tt.out
			got := Day08Part2(tt.in)
			if want != got {
				t.Fatalf("%q: want %d but got %d", string(tt.in),
					want, got)
			}
		})
	}
	want := 6
	buf := []byte{
		'"',
		'"',
	}
	got := Day08Part2(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2(t *testing.T) {
	f, err := os.Open(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	mm := 0
	nn := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Bytes()
		mm += len(line)
		nn += Day08Part2(line)
	}
	want := 2117
	got := nn - mm
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	buf, err := inputFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day08Part1(buf)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	buf, err := inputFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day08Part2(buf)
	}
}
