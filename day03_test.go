package adventofcode2015

import (
	"os"
	"testing"
)

var day3ExamplesPart1 = []struct {
	in  string
	out uint
}{
	{">", 2},
	{"^>v<", 4},
	{"^v^v^v^v^v", 2},
}

func TestDay3ExamplesPart1(t *testing.T) {
	for _, tt := range day3ExamplesPart1 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day3Part1([]byte(tt.in))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func inputDay3() (string, error) {
	buf, err := os.ReadFile(filename(3))
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func TestDay3Part1(t *testing.T) {
	const want = 2081
	s, err := inputDay3()
	if err != nil {
		t.Fatal(err)
	}
	got := Day3Part1([]byte(s))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

var day3ExamplesPart2 = []struct {
	in  string
	out uint
}{
	{"^>", 3},
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
}

func TestDay3ExamplesPart2(t *testing.T) {
	for _, tt := range day3ExamplesPart2 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Day3Part2([]byte(tt.in))
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay3Part2(t *testing.T) {
	const want = 2341
	s, err := inputDay3()
	if err != nil {
		t.Fatal(err)
	}
	got := Day3Part2([]byte(s))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
