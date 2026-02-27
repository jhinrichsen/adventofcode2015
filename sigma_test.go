package adventofcode2015

import "testing"

const sampleSigma = 13

func BenchmarkSigma(b *testing.B) {
	for b.Loop() {
		Sigma(sampleSigma)
	}
}

func TestSigma(t *testing.T) {
	for i := 1; i < len(A000203Seq); i++ {
		want := A000203Seq[i]
		got := Sigma(uint(i))
		if want != got {
			t.Fatalf("Sigma(%d) want %d but got %d", i, want, got)
		}
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
