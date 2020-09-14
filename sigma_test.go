package adventofcode2015

import (
	"fmt"
	"testing"
)

const sampleSigma = 13

func BenchmarkSigma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sigma(sampleSigma)
	}
}

func TestSigma(t *testing.T) {
	testSigmaFunction(t, Sigma)
}

// TODO broke
func testAnotherSigma(t *testing.T) {
	testSigmaFunction(t, AnotherSigma)
}

func testA000203(t *testing.T) {
	ch := make(chan uint, 2)
	go a000203(ch)
	// for i := 1; i < len(A000203Seq); i++ {
	for i := 1; i < len(A000203Seq); i++ {
		want := A000203Seq[i]
		got := <-ch
		if want != got {
			t.Fatalf("N=%d: want %d but got %d", i, want, got)
		}
	}
}

// TODO broke
func testSigmaGenerator(t *testing.T) {
	yield := sigmaGenerator()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sigma(%d) = %d\n", i, yield())
	}
}

// TODO broke
func exampleSigmaGenerator() {
	yield := sigmaGenerator()
	yield()              // n = 1
	yield()              // n = 2
	fmt.Println(yield()) // n = 3
	// Output: 4
}

// TODO broke
func testSigmaFunction(t *testing.T, f func(uint) uint) {
	for i := 1; i < len(A000203Seq); i++ {
		want := A000203Seq[i]
		got := f(uint(i))
		if want != got {
			t.Fatalf("N=%d: want %d but got %d", i, want, got)
		}
	}
}

// BenchmarkYieldAlt runs in about 2.8 ns/op.
func BenchmarkYieldAlt(b *testing.B) {
	gen := yieldAlt()
	for i := 0; i < b.N; i++ {
		gen()
	}
}

func TestYieldAlt(t *testing.T) {
	wants := []uint{1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6}
	alt := yieldAlt()
	for i, want := range wants {
		got := alt()
		if want != got {
			t.Fatalf("position %d: want %d but got %d",
				i, want, got)
		}
	}
}

func TestYieldIndex(t *testing.T) {
	wants := []uint{1, 2, 5, 7, 12, 15, 22, 26, 35, 40, 51, 57, 70, 77, 92}
	gen := yieldIndex()
	for i, want := range wants {
		got := gen()
		if want != got {
			t.Fatalf("position %d: want %d but got %d",
				i, want, got)
		}
	}
}

func BenchmarkSigmaRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SigmaRecursive(sampleSigma)
	}
}

func TestSigmaRecursive(t *testing.T) {
	// this one is dead slow, only test a subset
	for i := 1; i < len(A000203Seq)/2; i++ {
		want := A000203Seq[i]
		got := SigmaRecursive(uint(i))
		if want != got {
			t.Fatalf("sigma(%d): want %d but got %d", i, want, got)
		}
	}
}

func BenchmarkSigmaMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SigmaMemoized(sampleSigma)
	}
}

func TestSigmaMemoized(t *testing.T) {
	for i := 1; i < len(A000203Seq); i++ {
		want := A000203Seq[i]
		got := SigmaMemoized(uint(i))
		if want != got {
			t.Fatalf("sigma(%d): want %d but got %d", i, want, got)
		}
	}
}
