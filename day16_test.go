package adventofcode2015

import "testing"

func BenchmarkDay16Part1(b *testing.B) {
	benchWithParser(b, 16, true, NewDay16, Day16)
}

func BenchmarkDay16Part2(b *testing.B) {
	benchWithParser(b, 16, false, NewDay16, Day16)
}

func TestDay16Parse(t *testing.T) {
	sue, err := day16ParseSue("Sue 475: trees: 2, cars: 7, akitas: 8")
	if err != nil {
		t.Fatal(err)
	}
	if sue.number != 475 {
		t.Fatalf("want %d but got %d", 475, sue.number)
	}
	if sue.props[day16Trees] != 2 || sue.props[day16Cars] != 7 || sue.props[day16Akitas] != 8 {
		t.Fatalf("unexpected parsed properties: %+v", sue.props)
	}
}

func TestDay16Part1(t *testing.T) {
	testWithParser(t, 16, filename, true, NewDay16, Day16, uint(373))
}

func TestDay16Part2(t *testing.T) {
	testWithParser(t, 16, filename, false, NewDay16, Day16, uint(260))
}
