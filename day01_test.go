package adventofcode2015

import "testing"

func TestDay01Examples(t *testing.T) {
	for _, tt := range []struct {
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
	} {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day01Branching([]byte(tt.in), true)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay01Part1(t *testing.T) {
	testSolver(t, 1, filename, true, Day01, 232)
}

func TestDay01Part1Branching(t *testing.T) {
	testSolver(t, 1, filename, true, Day01Branching, 232)
}

func TestDay01Part1Branchless(t *testing.T) {
	testSolver(t, 1, filename, true, Day01Branchless, 232)
}

func TestDay01Part2Examples(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{
		{")", 1},
		{"()())", 5},
	} {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day01Branching([]byte(tt.in), false)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay01Part2(t *testing.T) {
	testSolver(t, 1, filename, false, Day01, 1783)
}

func TestDay01Part2Branching(t *testing.T) {
	testSolver(t, 1, filename, false, Day01Branching, 1783)
}

func TestDay01Part2Branchless(t *testing.T) {
	testSolver(t, 1, filename, false, Day01Branchless, 1783)
}

func BenchmarkDay01Part1(b *testing.B) {
	bench(b, 1, true, Day01Branching)
}

func BenchmarkDay01Part1Branchless(b *testing.B) {
	bench(b, 1, true, Day01Branchless)
}

func BenchmarkDay01Part2(b *testing.B) {
	bench(b, 1, false, Day01Branching)
}

func BenchmarkDay01Part2Branchless(b *testing.B) {
	bench(b, 1, false, Day01Branchless)
}
