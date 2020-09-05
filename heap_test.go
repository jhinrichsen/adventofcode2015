package adventofcode2015

import (
	"testing"
)

func TestCities(t *testing.T) {
	cities := []string{"london", "dublin", "edinburgh"}
	ch := make(chan []string)
	go heap(3, cities, ch)
	var perms [][]string
	for perm := range ch {
		perms = append(perms, perm)
	}
	want := Fac(uint(len(cities)))
	got := uint(len(perms))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
