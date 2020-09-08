package adventofcode2015

import (
	"fmt"
	"testing"
)

var examples = []struct {
	s, old, new string
	n           int
	result      string
}{
	{"A", "A", "B", 1, "B"},
	{"AABACA", "A", "B", 3, "AABBCA"},
}

func TestReplace(t *testing.T) {
	for i, tt := range examples {
		id := fmt.Sprintf("%d", i)
		t.Run(id, func(t *testing.T) {
			want := tt.result
			got := ReplaceNth(tt.s, tt.old, tt.new, tt.n)
			if want != got {
				t.Fatalf("want %q but got %q", want, got)
			}
		})
	}
}
