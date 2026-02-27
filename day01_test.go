package adventofcode2015

import "testing"

func TestDay1Examples(t *testing.T) {
	day1ExamplesPart1 := []struct {
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
	for _, tt := range day1ExamplesPart1 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day01([]byte(tt.in), true)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	bench(b, 1, true, Day01)
}

func BenchmarkDay01Part1Branchless(b *testing.B) {
	bench(b, 1, true, Day01Branchless)
}

func TestDay1Part1(t *testing.T) {
	testSolver(t, 1, filename, true, Day01, 232)
}

func TestDay1Part2Examples(t *testing.T) {
	day1ExamplesPart2 := []struct {
		in  string
		out int
	}{
		{")", 1},
		{"()())", 5},
	}
	for _, tt := range day1ExamplesPart2 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day01([]byte(tt.in), false)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	bench(b, 1, false, Day01)
}

func TestDay1Part2(t *testing.T) {
	testSolver(t, 1, filename, false, Day01, 1783)
}
