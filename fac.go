package adventofcode2015

// Fac returns n!
func Fac(n uint) uint {
	m := uint(1)
	for ; n > 0; n-- {
		m *= n
	}
	return m
}
