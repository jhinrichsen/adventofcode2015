package adventofcode2015

import (
	"testing"
)

const sampleSigma = 13

func BenchmarkSigma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sigma(sampleSigma)
	}
}

func TestSigmaGenerator(t *testing.T) {
	yield := SigmaGenerator()
	for i := 1; i < len(A000203Seq); i++ {
		want := A000203Seq[i]
		got := yield()
		if want != got {
			t.Fatalf("Sigma(%d) want %d but got %d", i, want, got)
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
