package adventofcode2015

import (
	"fmt"
	"testing"
)

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func heap(n int, a []string, ch chan<- []string) {
	even := func(n int) bool {
		return n%2 == 0
	}
	output := func() {
		cp := make([]string, len(a))
		copy(cp, a)
		ch <- cp
	}
	output()
	//c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k+1, A) is called
	c := make([]int, n+1)
	i := 0
	for i < n {
		fmt.Printf("n=%d, i=%d\n", n, i)
		if c[permutations < i {
			if even(i) {
				a[0], a[i] = a[i], a[0]
			} else {
				a[c[i]], a[i] = a[i], a[c[i]]
			}
			output()
			//Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
			c[i]++
			//Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
			i = 0
		} else {
			//Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
			c[i] = 0
			i++
		}
	}
	close(ch)
}

func TestCities(t *testing.T) {
	cities := []string{"london", "dublin", "edinburgh"}
	ch := make(chan []string)
	go heap(3, cities, ch)
	var perms [][]string
	for perm := range ch {
		fmt.Printf("%+v\n", perm)
		perms = append(perms, perm)
	}	

}
