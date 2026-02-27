package adventofcode2015

import (
	"strings"
	"testing"
)

func BenchmarkDay18Part1(b *testing.B) {
	benchWithParser(b, 18, true, NewDay18, Day18)
}

func BenchmarkDay18Part2(b *testing.B) {
	benchWithParser(b, 18, false, NewDay18, Day18)
}

func TestDay18EvolutionExample(t *testing.T) {
	steps := []string{
		`.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
		`..##..
..##.#
...##.
......
#.....
#.##..`,
		`..###.
......
..###.
......
.#....
.#....`,
		`...#..
......
...#..
..##..
......
......`,
		`......
......
..##..
..##..
......
......`,
	}
	lines := strings.Split(steps[0], "\n")
	puzzle, err := NewDay18(lines)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, len(puzzle.buf))
	copy(buf, puzzle.buf)
	for i := 1; i < len(steps); i++ {
		buf = day18Step(buf, puzzle.w, puzzle.h)
		if got := day18String(buf, puzzle.w, puzzle.h); got != steps[i] {
			t.Fatalf("step %d mismatch\nwant:\n%s\n\ngot:\n%s", i, steps[i], got)
		}
	}
	if got := day18CountOn(buf); got != 4 {
		t.Fatalf("want %d but got %d", 4, got)
	}
}

func TestDay18Part2Example(t *testing.T) {
	lines := strings.Split(`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, "\n")
	puzzle, err := NewDay18(lines)
	if err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, len(puzzle.buf))
	copy(buf, puzzle.buf)
	day18SetCorners(buf, puzzle.w, puzzle.h, day18LightOn)
	for range 5 {
		buf = day18Step(buf, puzzle.w, puzzle.h)
		day18SetCorners(buf, puzzle.w, puzzle.h, day18LightOn)
	}
	if got := day18CountOn(buf); got != 17 {
		t.Fatalf("want %d but got %d", 17, got)
	}
}

func TestDay18Part1(t *testing.T) {
	testWithParser(t, 18, filename, true, NewDay18, Day18, uint(1_061))
}

func TestDay18Part2(t *testing.T) {
	testWithParser(t, 18, filename, false, NewDay18, Day18, uint(1_006))
}

func day18String(buf []byte, w, h int) string {
	var sb strings.Builder
	for y := 0; y < h; y++ {
		sb.Write(buf[y*w : (y+1)*w])
		if y < h-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

