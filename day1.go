package adventofcode2015

const (
	openPar  = '('
	closePar = ')'
)

// Day1Part1 returns number of opening braces minus number of closing braces.
func Day1Part1(buf []byte) int {
	floor := 0
	for _, b := range buf {
		if b == openPar {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

// Day1Part2 returns position where floor gets negative.
func Day1Part2(buf []byte) uint {
	var floor int
	for i, b := range buf {
		if b == openPar {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			// position is 1-based
			return uint(i) + 1
		}
	}
	return 0
}
