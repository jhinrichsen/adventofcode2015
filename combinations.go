//  Restricted integer combinations after Knuth, TAOCP pre-fascicle 3b, 7.2.1.4
// Algorith H, 2004, after C. H. Hindenburg, Infinitinomii Dignitatum Exponentis
// Indeterminati, 1789

package adventofcode2015

// AlgorithmH partitions into m parts. Writes digits of integer tuples into
// channel. Reference: Knuth TAOCP, fascicle 3b, Chap 7.2.1.4, Algorith H. Dates
// back to C.F. Hindenburg (hence H) in his essay "Infinitinomii Dignitatum
// Exponentis Indeterminati".
// There are competing, more generic algorithms such as
// "A Unified Approach to Algorithms Generating Unrestricted and Restricted
// Integer Compositions and Integer Partitions" J. D. OPDYKE, J. Math. Modelling
// and Algorithms (2009) V9 N1, p.53 - 97, in the end we don't need more
// features, and nobody ever got fired for picking Knuth.
func AlgorithmH(n, m int, c chan<- []int) {
	// algorith uses 1-based arrays, so a[0] is unused
	a := make([]int, m+2)
	// H1 [Initialize]
	a[1] = n - m + 1
	for j := 2; j <= m; j++ {
		a[j] = 1
	}
	a[m+1] = -1
H2:
	// H2 [Visit]
	// Implementation notice: slices cannot be passed via channels without
	// being overwritten constantly.
	cp := make([]int, m)
	copy(cp, a[1:m+1])
	c <- cp
	if a[2] >= a[1]-1 {
		goto H4
	}
	// H3 [Tweak a₁ and a₂]
	a[1] = a[1] - 1
	a[2] = a[2] + 1
	goto H2
H4:
	// H4 [Find j]
	j := 3
	s := a[1] + a[2] - 1
	for a[j] >= a[1]-1 {
		s = s + a[j]
		j++
	}
	// H5 [Increase aj]
	if j > m {
		close(c)
		return
	}
	x := a[j] + 1
	a[j] = x
	j--
	// H6 [Tweak a₁...aj ]
	for j > 1 {
		a[j] = x
		s = s - x
		j = j - 1
	}
	a[1] = s
	goto H2
}

// AlgorithmT implements Knuth Fascicles 3a, 7.2.1.3, lexicographic
// combinations.
func AlgorithmT(t, n int, ch chan<- []int) {
	// algorithm is 1-based
	c := make([]int, t+3)
	var x int

	// T1 [Initialize]
	for j := 1; j <= t; j++ {
		c[j] = j - 1
	}
	c[t+1] = n
	c[t+2] = 0
	j := t
T2:
	// T2 [Visit]
	cp := make([]int, t+1)
	copy(cp, c)
	ch <- cp
	if j > 0 {
		x = j
		goto T6
	}

	// T3 [Easy case?]
	if c[1]+1 < c[2] {
		c[1]++
		goto T2
	}
	j = 2

	// T4 [Find j]
	c[j-1] = j - 2
	x = c[j] + 1
	for x == c[j+1] {
		j++
	}

	// T5 [Done?]
	if j > t {
		close(ch)
		return
	}
T6:
	// T6 [Increase cj]
	c[j] = x
	j--
	goto T2
}

// KCompositions returns all combinations of k digits that add up to n.
func KCompositions(n, k int, ch chan<- []int) {
	// TODO
	// panic(k > n)
	// panic(n == 0)
	// panic(k == 0)

	v := make([]int, k)
	visit := func() {
		// visit
		cp := make([]int, k)
		copy(cp, v)
		ch <- cp

	}
	var j int
	v[0] = n - k + 1
	for j = 1; j < k; j++ {
		v[j] = 1
	}
	visit()

	for k > 1 {
		if v[0] > 1 {
			v[0]--
			v[1]++
		} else {
			for j = 1; j < k; j++ {
				if v[j] > 1 {
					break
				}
			}
			if j >= k-1 {
				close(ch)
				return
			}
			v[j+1]++
			v[0] = v[j] - 1
			v[j] = 1
		}
		visit()
	}
}
