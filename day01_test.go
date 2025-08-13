package adventofcode2015

import (
	"os"
	"testing"
)

var day1ExamplesPart1 = []struct {
	in  string
	out int
}{
	{"(())", 0},
	{"()()", 0},
	{"(((", 3},
	{"(()(()(", 3},
	{"))(((((", 3},
	{"())", -1},
	{"))(", -1},
	{")))", -3},
	{")())())", -3},
}

func day1Input() ([]byte, error) {
	return os.ReadFile(filename(1))
}

func TestDay1Examples(t *testing.T) {
	for _, tt := range day1ExamplesPart1 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day1Part1([]byte(tt.in))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay1Part1(b *testing.B) {
	buf, err := day1Input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day1Part1(buf)
	}
}

func BenchmarkDay1Part1Branchless(b *testing.B) {
	buf, err := day1Input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day1Part1Branchless(buf)
	}
}

func TestDay1Part1(t *testing.T) {
	const want = 232
	buf, err := day1Input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day1Part1(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

var day1ExamplesPart2 = []struct {
	in  string
	out int
}{
	{")", 1},
	{"()())", 5},
}

func TestDay1Part2Examples(t *testing.T) {
	for _, tt := range day1ExamplesPart2 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day1Part2([]byte(tt.in))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	buf, err := day1Input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day1Part2(buf)
	}
}

func TestDay1Part2(t *testing.T) {
	const want = 1783
	buf, err := day1Input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day1Part2(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
