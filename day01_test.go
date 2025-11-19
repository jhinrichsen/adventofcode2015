package adventofcode2015

import (
	"os"
	"testing"
)

var day01ExamplesPart1 = []struct {
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

func day01Input() ([]byte, error) {
	return os.ReadFile(filename(1))
}

func TestDay01Examples(t *testing.T) {
	for _, tt := range day01ExamplesPart1 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day01Part1([]byte(tt.in))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	buf, err := day01Input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day01Part1(buf)
	}
}

func BenchmarkDay01Part1Branchless(b *testing.B) {
	buf, err := day01Input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day01Part1Branchless(buf)
	}
}

func TestDay01Part1(t *testing.T) {
	const want = 232
	buf, err := day01Input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day01Part1(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

var day01ExamplesPart2 = []struct {
	in  string
	out int
}{
	{")", 1},
	{"()())", 5},
}

func TestDay01Part2Examples(t *testing.T) {
	for _, tt := range day01ExamplesPart2 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day01Part2([]byte(tt.in))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	buf, err := day01Input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day01Part2(buf)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 1783
	buf, err := day01Input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day01Part2(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
