package adventofcode2015

import "testing"

func TestDay04Part1Example1(t *testing.T) {
	testSolver(t, 4, example1Filename, true, Day04, uint(609_043))
}

func TestDay04Part1Example2(t *testing.T) {
	testSolver(t, 4, example2Filename, true, Day04, uint(1_048_970))
}

func BenchmarkDay04Part1(b *testing.B) {
	bench(b, 4, true, Day04)
}

func TestDay04Part1(t *testing.T) {
	testSolver(t, 4, filename, true, Day04, uint(254_575))
}

func BenchmarkDay04Part2(b *testing.B) {
	bench(b, 4, false, Day04)
}

func TestDay04Part2(t *testing.T) {
	testSolver(t, 4, filename, false, Day04, uint(1_038_736))
}
