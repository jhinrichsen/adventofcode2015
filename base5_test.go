package adventofcode2015

import (
	"reflect"
	"testing"
)

func TestBase5Inc(t *testing.T) {
	want := []byte{0, 1, 0}
	b := NewBase5(3)
	b.Inc()
	b.Inc()
	b.Inc()
	b.Inc()
	b.Inc()
	got := b.Buf
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}
