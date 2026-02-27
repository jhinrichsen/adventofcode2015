package adventofcode2015

import "testing"

func BenchmarkDay12Part1(b *testing.B) {
	bench(b, 12, true, Day12)
}

func BenchmarkDay12Part2(b *testing.B) {
	bench(b, 12, false, Day12)
}

func TestDay12SamplesPart1(t *testing.T) {
	tests := []struct {
		in   string
		want uint
	}{
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
		{`{}`, 0},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Day12([]byte(tt.in), true)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("%q: want %d but got %d", tt.in, tt.want, got)
			}
		})
	}
}

func TestDay12SamplesPart2(t *testing.T) {
	tests := []struct {
		in   string
		want uint
	}{
		{`[1,2,3]`, 6},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Day12([]byte(tt.in), false)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("%q: want %d but got %d", tt.in, tt.want, got)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	testSolver(t, 12, filename, true, Day12, uint(111_754))
}

func TestDay12Part2(t *testing.T) {
	testSolver(t, 12, filename, false, Day12, uint(65_402))
}
