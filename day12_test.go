package adventofcode2015

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

var day11Samples = []struct {
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

func TestDay12Samples(t *testing.T) {
	for _, tt := range day11Samples {
		id := tt.in
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := sum(tt.in)
			if want != got {
				t.Fatalf("%q: want %d but got %d", id, want, got)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	filename := "testdata/day12.txt"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	const want = 111754
	got := sum(string(buf))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
