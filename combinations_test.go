package adventofcode2015

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestAlgorithmH(t *testing.T) {
	want := []int{
		8111,
		7211,
		6311,
		5411,
		6221,
		5321,
		4421,
		4331,
		5222,
		4322,
		3332,
	}
	c := make(chan []int)
	var got []int
	go AlgorithmH(11, 4, c)
	got, err := read(c)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestAlgorithmT(t *testing.T) {
	c := make(chan []int)
	go AlgorithmT(2, 4, c)
	_, err := read(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKCompositions(t *testing.T) {
	want := []int{
		31111,
		22111,
		13111,
		21211,
		12211,
		11311,
		21121,
		12121,
		11221,
		11131,
		21112,
		12112,
		11212,
		11122,
		11113,
	}
	c := make(chan []int)
	go KCompositions(7, 5, c)
	got, err := read(c)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func read(ch chan []int) ([]int, error) {
	var is []int
	for digits := range ch {
		var sb strings.Builder
		for _, digit := range digits {
			sb.WriteByte('0' + byte(digit))
		}
		n, err := strconv.Atoi(sb.String())
		if err != nil {
			return is, err
		}
		is = append(is, n)
	}
	return is, nil
}
