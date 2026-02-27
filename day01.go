package adventofcode2015

// Day01 solves day 1 for the selected part.
func Day01(buf []byte, part1 bool) (int, error) {
	floor := 0
	for i, b := range buf {
		if b == '(' {
			floor++
		} else {
			floor--
		}
		if !part1 && floor < 0 {
			// position is 1-based
			return i + 1, nil
		}
	}
	if part1 {
		return floor, nil
	}
	return 0, nil
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
		// floor += int('(' + ')' - b - b)
		// floor += int('(' + ')' - b - b)
		floor += int(81 - 2*b)
	}
	return
}

// Day01Branchless solves day 1 using the branchless part 1 variant.
func Day01Branchless(buf []byte, part1 bool) (int, error) {
	if part1 {
		return Day1Part1Branchless(buf), nil
	}
	return Day01(buf, false)
}
