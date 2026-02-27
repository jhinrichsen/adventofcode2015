package adventofcode2015

import "testing"

func BenchmarkDay19Part1(b *testing.B) {
	benchWithParser(b, 19, true, NewDay19, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchWithParser(b, 19, false, NewDay19, Day19)
}

func TestDay19ReplaceNth(t *testing.T) {
	tests := []struct {
		s, old, new string
		n           int
		want        string
	}{
		{"A", "A", "B", 1, "B"},
		{"AABACA", "A", "B", 3, "AABBCA"},
	}
	for _, tt := range tests {
		if got := replaceNth(tt.s, tt.old, tt.new, tt.n); got != tt.want {
			t.Fatalf("want %q but got %q", tt.want, got)
		}
	}
}

func TestDay19Part1Example(t *testing.T) {
	puzzle, err := NewDay19([]string{
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOH",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := Day19(puzzle, true); got != 4 {
		t.Fatalf("want %d but got %d", 4, got)
	}
}

func TestDay19Part2Example1(t *testing.T) {
	puzzle, err := NewDay19([]string{
		"e => H",
		"e => O",
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOH",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := Day19(puzzle, false); got != 3 {
		t.Fatalf("want %d but got %d", 3, got)
	}
}

func TestDay19Part2Example2(t *testing.T) {
	puzzle, err := NewDay19([]string{
		"e => H",
		"e => O",
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOHOHO",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := Day19(puzzle, false); got != 6 {
		t.Fatalf("want %d but got %d", 6, got)
	}
}

func TestDay19Part1(t *testing.T) {
	testWithParser(t, 19, filename, true, NewDay19, Day19, uint(576))
}

func TestDay19Part2(t *testing.T) {
	testWithParser(t, 19, filename, false, NewDay19, Day19, uint(207))
}
