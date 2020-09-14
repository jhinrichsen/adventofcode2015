package adventofcode2015

import (
	"fmt"
	"math"
)

// A000203Seq is the beginning sequence of OEIS A000203, sum of
// divisors.
// Source: "Découverte d'une loi tout extraordinaire des nombres, par rapport à
// la somme de leurs diviseurs.", Leonhard Euler, Berlin, 22. Juni 1747.
var A000203Seq = [...]uint{
	0,
	1,  // 1
	3,  // 2
	4,  // 3
	7,  // 4
	6,  // 5
	12, // 6
	8,  // 7
	15, // 8
	13, // 9
	18, // 10
	12, // 11
	28, // 12
	14, // 13
	24, // 14
	24, // 15
	31, // 16
	18, // 17
	39, // 18
	20, // 19
	42, // 20
	32, // 21
	36, // 22
	24, // 23
	60, // 24
	31, // 25
	42, // 26
	40, // 27
	56, // 28
	30, // 29
	72, // 30
}

// sigmaGenerator generates sum of divisors, one at a time.
// It is not exported yet, because the current algorithm needs a final N, and
// cannot be used for generating the next from the previous.
func sigmaGenerator() func() uint {
	var n, sum uint
	n = 1
	return func() uint {
		_, frac := math.Modf(float64(n) / float64(n))
		f := math.Mod(Sgnf(frac)+1.0, 2)
		p := n * uint(f)
		sum += p
		n++
		return sum
	}
}

// Sigma function, sum of divisors σ(n).
// Divisors are also known as factors.
// Sum of its divisors is also known as the Sigma function.
// Sigma is OEIS sequence https://oeis.org/A000203.
func Sigma(n uint) (sum uint) {
	if n == 0 {
		return 0
	}
	for i := uint(1); i <= n; i++ {
		_, frac := math.Modf(float64(n) / float64(i))
		f := math.Mod(Sgnf(frac)+1.0, 2)
		p := i * uint(f)
		sum += p
	}
	return
}

// a000203 function, sum of divisors σ(n).
// Generates the a000203 sequence into c, starting at a000203(1) = 1.
// TODO work in progres, generator not prime yet.
func a000203(c chan<- uint) {
	yield := sigmaGenerator()
	for {
		c <- yield()
	}
}

// Sgnf returns the signature of a float value.
// https://en.wikipedia.org/wiki/Sign_function
// https://github.com/golang/go/issues/3743
// Results, including some IEEE754 edge cases:
// Sgnf(f > 0.0) = 1.0
// Sgnf(f < 0.0) = -1.0
// Sgnf(0.0) = 0.0
// Sgnf(-0.0) = 0.0
// Sgnf(NaN) = NaN
// Sgnf(Inf) = 1.0
// Sgnf(-Inf) = -1.0
func Sgnf(f float64) float64 {
	if f > 0.0 {
		return 1.0
	}
	if f < 0.0 {
		return -1.0
	}
	if f == 0.0 {
		return 0.0
	}
	// -0: some trickery to not fall into const/ sign issues
	// https://github.com/golang/go/issues/2196
	if f == math.Copysign(0, -1) {
		return 0.0
	}
	if math.IsNaN(f) {
		return math.NaN()
	}
	if f == math.Inf(1) {
		return 1.0
	}
	if f == math.Inf(-1) {
		return -1.0
	}
	panic(fmt.Sprintf("unknown float value %+v", f))
}

// Sgn returns the signature of a value.
// https://en.wikipedia.org/wiki/Sign_function
// https://github.com/golang/go/issues/3743
func Sgn(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	panic(fmt.Sprintf("cannot determine sign of %d", n))
}

// AnotherSigma implements https://www.xarg.org/2016/06/calculate-the-sum-of-divisors/.
// TODO broke
func AnotherSigma(n uint) uint {
	var sum uint
	for i := uint(1); i <= n; i++ {
		p := math.Mod(float64(n), float64(i))
		sum += uint(p)
	}
	return n*n - sum
}

// I. In the alternation of the signs + and −, each repeats two at a
// time.
func yieldSign() func() int {
	signs := []int{1, 1, -1, -1}
	idx := 0
	return func() int {
		s := signs[idx]
		idx++
		if idx == len(signs) {
			idx = 0
		}
		return s
	}
}

// all the natural numbers, 1, 2, 3, 4, 5, 6
func yieldN() func() uint {
	var n uint
	return func() uint {
		n++
		return n
	}
}

// odd numbers 3, 5, 7, 9, 11
func yieldOdd() func() uint {
	n := uint(1)
	return func() uint {
		n++
		n++
		return n
	}
}

// yieldAt iterates natural numbers alternating with odd numbers.
// Diff.1,3,2,5, 3, 7, 4, 9, 5, 11,6, 13,7, 15,8...
func yieldAlt() func() uint {
	var b bool
	n := yieldN()
	odd := yieldOdd()
	return func() uint {
		b = !b
		if b {
			return n()
		}
		return odd()
	}
}

// yieldIndex returns N. 1, 2, 5, 7, 12, 15, 22, 26, 35, 40, 51, 57, 70, 77,...
func yieldIndex() func() uint {
	n := uint(1)
	diff := yieldAlt()
	return func() uint {
		previous := n
		n += diff()
		return previous
	}
}

// SigmaRecursiveF accepts a recursive function to allow for optional
// memozation.
func SigmaRecursiveF(n uint, f func(uint) uint) (sum uint) {
	genIdx := yieldIndex()
	genSign := yieldSign()

	for {
		idx := genIdx()
		// III. Although this series goes to infinity, we only have to
		// take, in each case, the terms starting where the number after
		// the σ sign is still positive, omitting those that contain
		// negative numbers.
		if n < idx {
			break
		}
		var sigma uint
		// IV. If it happens that the term σ(0) appears in this formula,
		// since its value is indeterminate in itself, we must, in each
		// case, instead of σ(0) put the given number itself.
		j := n - idx
		if j == 0 {
			sigma = n
		} else {
			sigma = f(j)
		}
		sign := genSign()
		if sign > 0 {
			sum += sigma
		} else {
			sum -= sigma
		}
	}
	return
}

// SigmaRecursive is an implementation of Leonhard Euler, "Discovery of a most
// extraordinary law of numbers§", Berlin, 22.06.1747, section 5.
func SigmaRecursive(n uint) (sum uint) {
	return SigmaRecursiveF(n, SigmaRecursive)
}

// sigmas is a momoization cache.
var sigmas map[uint]uint

// initialize sigmas cache
func init() {
	sigmas = make(map[uint]uint)
}

// SigmaMemoized implements a memoized version of SigmaRecursive.
func SigmaMemoized(n uint) uint {
	if cached, ok := sigmas[n]; ok {
		return cached
	}
	sigma := SigmaRecursiveF(n, SigmaMemoized)
	sigmas[n] = sigma
	return sigma
}
