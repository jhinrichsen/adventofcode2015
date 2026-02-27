package adventofcode2015

import "testing"

func TestDay03ExamplesPart1(t *testing.T) {
	day3ExamplesPart1 := []struct {
		in  string
		out uint
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}
	for _, tt := range day3ExamplesPart1 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day03([]byte(tt.in), true)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	bench(b, 3, true, Day03)
}

func TestDay03Part1(t *testing.T) {
	testSolver(t, 3, filename, true, Day03, uint(2081))
}

func TestDay03ExamplesPart2(t *testing.T) {
	day3ExamplesPart2 := []struct {
		in  string
		out uint
	}{
		{"^>", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for _, tt := range day3ExamplesPart2 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day03([]byte(tt.in), false)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	bench(b, 3, false, Day03)
}

func TestDay03Part2(t *testing.T) {
	testSolver(t, 3, filename, false, Day03, uint(2341))
}
