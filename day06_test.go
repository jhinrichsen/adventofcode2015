package adventofcode2015

import "testing"

func BenchmarkDay06Part1(b *testing.B) {
	benchLinesErr(b, 6, true, func(lines []string, part1 bool) (uint, error) {
		if part1 {
			return Day6Part1(lines)
		}
		return Day6Part2(lines)
	})
}

func BenchmarkDay06Part2(b *testing.B) {
	benchLinesErr(b, 6, false, func(lines []string, part1 bool) (uint, error) {
		if part1 {
			return Day6Part1(lines)
		}
		return Day6Part2(lines)
	})
}

func TestDay6Part1(t *testing.T) {
	const want = 400_410
	lines := linesFromFilename(t, filename(6))
	got, err := Day6Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part2(t *testing.T) {
	const want = 15_343_601
	lines := linesFromFilename(t, filename(6))
	got, err := Day6Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
