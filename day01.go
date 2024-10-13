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

// Day1Part1Branchless returns number of opening braces minus number of closing
// braces.
func Day1Part1Branchless(buf []byte) (floor int) {
	for _, b := range buf {
		/*
			floor += int(closePar - b) // 40 -> 1, 41 -> 0
			floor -= int(b - 40)       // 40 -> 0, 41 -> 1
		*/
		// floor += int((closePar - b) - (b - 40))
		// floor += int((closePar - b) -b + 40)
		// floor += int(closePar - b - b + 40)
		// floor += int(openPar + closePar - b - b)
		// floor += int(openPar + closePar - b - b)
		floor += int(81 - 2*b)
	}
	return
}

// Day1Part2 returns position where floor gets negative.
func Day1Part2(buf []byte) int {
	var floor int
	for i, b := range buf {
		if b == openPar {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			// position is 1-based
			return i + 1
		}
	}
	return 0
}
