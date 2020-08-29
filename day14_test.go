package adventofcode2015

import (
	"fmt"
	"testing"
)

func TestDay14ExampleParsing(t *testing.T) {
	filename := exampleFilename(14)
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("want %d but got %d", 2, len(lines))
	}
	var rs []Reindeer
	for i, line := range lines {
		r, err := newReindeer(line)
		if err != nil {
			t.Fatalf("%s col %d: %+v", filename, i, err)
		}
		rs = append(rs, r)
	}
	want0 := Reindeer{"Comet", 14, 10, 127}
	got0 := rs[0]
	if want0 != got0 {
		t.Fatalf("want %+v but got %+v", want0, got0)
	}
}

func TestDay14Example(t *testing.T) {
	filename := exampleFilename(14)
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	var rs []Reindeer
	for i, line := range lines {
		r, err := newReindeer(line)
		if err != nil {
			t.Fatalf("%s col %d: %+v", filename, i, err)
		}
		rs = append(rs, r)
	}
	comet := rs[0]
	dancer := rs[1]
	var distances = []struct {
		r       Reindeer
		seconds uint
		km      uint
	}{
		{comet, 1, 14},
		{dancer, 1, 16},
		{comet, 10, 140},
		{dancer, 10, 160},
		{comet, 11, 140},
		{dancer, 11, 176},
		{comet, 1000, 1120},
		{dancer, 1000, 1056},
	}
	for _, tt := range distances {
		id := fmt.Sprintf("%+v", tt)
		t.Run(id, func(t *testing.T) {
			want := tt.km
			got := tt.r.km(tt.seconds)
			if want != got {
				t.Fatalf("%s after %d sec: want %d but got %d",
					tt.r.Name, tt.seconds, want, got)
			}
		})
	}
}

func TestDay14Part1(t *testing.T) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		t.Fatal(err)
	}
	var rs []Reindeer
	for i, line := range lines {
		r, err := newReindeer(line)
		if err != nil {
			t.Fatalf("file %q, line %d: %+v",
				filename(14), i, err)
		}
		rs = append(rs, r)
	}
	const want = 2655
	got := Day14Part1(rs, 2503)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part2Example(t *testing.T) {
	filename := exampleFilename(14)
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	var rs []Reindeer
	for i, line := range lines {
		r, err := newReindeer(line)
		if err != nil {
			t.Fatalf("%s col %d: %+v", filename, i, err)
		}
		rs = append(rs, r)
	}

	m := make(map[string]ReindeerScore)
	for _, sc := range scores(rs, 1000) {
		m[sc.Name] = sc
	}
	const wantDancer = 689
	gotDancer := m["Dancer"].Score
	if wantDancer != gotDancer {
		t.Fatalf("want %d but got %d", wantDancer, gotDancer)
	}
	const wantComet = 312
	gotComet := m["Comet"].Score
	if wantComet != gotComet {
		t.Fatalf("want %d but got %d", wantComet, gotComet)
	}
}

func TestDay14Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		t.Fatal(err)
	}
	var rs []Reindeer
	for i, line := range lines {
		r, err := newReindeer(line)
		if err != nil {
			t.Fatalf("file %q, line %d: %+v",
				filename(14), i, err)
		}
		rs = append(rs, r)
	}
	const want = 1059
	got := Day14Part2(rs, 2503)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
