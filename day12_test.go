package adventofcode2015

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var day12SamplesPart1 = []struct {
	in  string
	out int
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

var day12SamplesPart2 = []struct {
	in  string
	out int
}{
	{`[1,2,3]`, 6},
	{`[1,{"c":"red","b":2},3]`, 4},
	{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
	{`[1,"red",5]`, 6},
}

// Neither strconv.Atoi() nor fmt.Sscanf() support recovery of illegal reads.
func NoTestParse(t *testing.T) {
	const s = "30a 4b -5cd"
	r := strings.NewReader(s)
	var i int
	n, err := fmt.Fscanf(r, "%d", &i)
	fmt.Printf("i=%d, n=%d, err=%+v\n", i, n, err)

	if n == 1 {
		l := len(strconv.Itoa(i))
		var bs = make([]byte, l)
		m, err := r.Read(bs)
		if err != nil {
			log.Fatal(err)
		}
		if m != l {
			log.Fatalf("Need to read %d bytes but read %d\n",
				l, m)
		}
	}
	n, err = fmt.Fscanf(r, "%d", &i)
	fmt.Printf("i=%d, n=%d, err=%+v\n", i, n, err)
}

func TestDay12SamplesPart1(t *testing.T) {
	for _, tt := range day12SamplesPart1 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := sum2(tt.in)
			if want != got {
				t.Fatalf("%q: want %d but got %d", id, want, got)
			}
		})
	}
}

func TestDay12Part1Sum1(t *testing.T) {
	filename := "testdata/day12.txt"
	buf, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	const want = 111754
	got := sum(string(buf))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part1Sum2(t *testing.T) {
	filename := "testdata/day12.txt"
	buf, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	const want = 65402
	got := sum2(string(buf))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12SamplesPart2(t *testing.T) {
	for _, tt := range day12SamplesPart2 {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := sum2(tt.in)
			if want != got {
				t.Fatalf("%q: want %d but got %d", id, want, got)
			}
		})
	}
}

// This is where we really need some parsing.
func TestDay12Part2(t *testing.T) {
	filename := "testdata/day12.txt"
	buf, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	// "The empty interface says nothing" says Rob. It says any says Ian.
	var data interface{}
	if err := json.Unmarshal(buf, &data); err != nil {
		t.Fatal(err)
	}
	want := 65402
	got := sum2(string(buf))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
