package adventofcode2015

import "testing"

func TestDay05Part1Example(t *testing.T) {
	examples := []struct {
		in   string
		want uint
	}{
		{"ugknbfddgicrmopn", 1},
		{"aaa", 1},
		{"jchzalrnumimnmhp", 0},
		{"haegwjzuvuyypxyu", 0},
		{"dvszwmarrgswjxmb", 0},
	}
	for _, tt := range examples {
		t.Run(tt.in, func(t *testing.T) {
			got := Day05([]string{tt.in}, true)
			if tt.want != got {
				t.Fatalf("want %v but got %v", tt.want, got)
			}
		})
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	benchLines(b, 5, true, Day05)
}

func TestDay05Part1(t *testing.T) {
	testLines(t, 5, filename, true, Day05, uint(238))
}

func TestDay05Part2Example(t *testing.T) {
	examples := []struct {
		in   string
		want uint
	}{
		{"qjhvhtzxzqqjkmpb", 1},
		{"xxyxx", 1},
		{"uurcxstgmygtbstg", 0},
		{"ieodomkazucvgmuy", 0},
	}
	for _, tt := range examples {
		t.Run(tt.in, func(t *testing.T) {
			got := Day05([]string{tt.in}, false)
			if tt.want != got {
				t.Fatalf("want %v but got %v", tt.want, got)
			}
		})
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	benchLines(b, 5, false, Day05)
}

func TestDay05Part2(t *testing.T) {
	testLines(t, 5, filename, false, Day05, uint(69))
}
