package adventofcode2015

import "math"

type day25 struct {
	n    uint // 1-based index
	code uint
}

func newDay25() day25 {
	return day25{1, 20151125}
}

// next sets next code by calculating x(n+1) = x(n) * 252533 % 33554393.
func (a *day25) next() {
	a.n++
	a.code = a.code * 252533 % 33554393
}

// x returns the x position of index n.
// x -> 1 1 2 1 2 3 1 2 3 4 1 2 3 4 5
// The sequence is known as fractal sequence, https://oeis.org/A002260
func (a day25) x() uint {
	m := uint(math.Floor((-1 + math.Sqrt(float64(8*a.n-7))) / 2))
	return a.n - m*(m+1)/2
}

// y returns the y position of index n.
// y -> 1 2 1 3 2 1 4 3 2 1
// The sequence is known as fractal sequence, https://oeis.org/A004736
func (a day25) y() uint {
	t := uint(math.Floor(-1+math.Sqrt(float64(8*a.n-7)))) / 2
	return (t*t+3*t+4)/2 - a.n
}
