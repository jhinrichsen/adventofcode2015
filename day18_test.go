package adventofcode2015

import (
	"io/ioutil"
	"testing"
)

var (
	exampleSteps = []string{
		`.#.#.#
...##.
#....#
..#...
#.#..#
####..`, `..##..
..##.#
...##.
......
#.....
#.##..`, `..###.
......
..###.
......
.#....
.#....`, `...#..
......
...#..
..##..
......
......`, `......
......
..##..
..##..
......
......`,
	}

	exampleStepsPart2 = []string{
		`##.#.#
...##.
#....#
..#...
#.#..#
####.#`, `#.##.#
####.#
...##.
......
#...#.
#.####`, `#..#.#
#....#
.#.##.
...##.
.#..##
##.###`, `#...##
####.#
..##.#
......
##....
####.#`, `#.####
#....#
...#..
.##...
#.....
#.#..#`, `##.###
.##..#
.##...
.##...
#.#...
##...#`,
	}
)

func TestOff(t *testing.T) {
	const want = true
	b := []byte(exampleSteps[0])[0]
	got := b == lightOff
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestOn(t *testing.T) {
	const want = true
	b := []byte(exampleSteps[0])[1]
	got := b == lightOn
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestString(t *testing.T) {
	want := exampleSteps[0]
	g, err := newGrid(want)
	if err != nil {
		t.Fatal(err)
	}
	got := g.String()
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}

}

func TestDay18Example(t *testing.T) {
	const want = 4
	g, err := newGrid(exampleSteps[0])
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i < len(exampleSteps); i++ {
		g.step()
		want := exampleSteps[i]
		got := g.String()
		if want != got {
			t.Fatalf("step %d: want \n%+v\n but got \n%v",
				i, want, got)
		}
	}
	got := g.on()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// hashes returns occurences of '#' in a file.
func hashes(filename string) (uint, error) {
	var n uint
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return n, err
	}
	for _, b := range buf {
		if b == '#' {
			n++
		}
	}
	return n, nil
}

func TestDay18ExampleOn(t *testing.T) {
	const want = 15
	g, err := newGrid(exampleSteps[0])
	if err != nil {
		t.Fatal(err)
	}
	got := g.on()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Line0(t *testing.T) {
	const want = "#..####.##..#...#..#...#...###.#.#.#..#....#.##..#...##...#..#.....##..#####....#.##..##....##.#...."
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	got := lines[0]
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}

}

// gridFromFilename is a helper to construct a grid from a file.
func gridFromFilename(filename string) (grid, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return grid{}, err
	}
	return newGrid(string(buf))
}

func TestDay18Line0On(t *testing.T) {
	const want = 41
	g, err := gridFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	var on, off uint
	for _, b := range g.buf[0] {
		if b == lightOn {
			on++
		} else if b == lightOff {
			off++
		} else {
			t.Fatalf("unknown state: %q", b)
		}
	}
	if on+off != 100 {
		t.Fatalf("want on+off=100 but got on=%d, off=%d", on, off)
	}
	got := on
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Number of # in input file must match on() after parsing.
func TestDay18On(t *testing.T) {
	want, err := hashes(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	g, err := gridFromFilename(filename(18))
	got := g.on()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 1061
	g, err := gridFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	x, y := g.dim()
	if x != 100 {
		t.Fatalf("want x=100 but got x=%d", x)
	}
	if y != 100 {
		t.Fatalf("want y=100 but got y=%d", y)
	}
	got := Day18Part1(g, 100)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2Example(t *testing.T) {
	const want = 17
	g, err := newGrid(exampleSteps[0])
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part2(g, 5)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = 1006
	g, err := gridFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part2(g, 100)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
