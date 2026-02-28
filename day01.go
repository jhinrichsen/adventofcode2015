package adventofcode2015

func Day01(buf []byte, part1 bool) (int, error) {
	return Day01Branchless(buf, part1)
}

func Day01Branching(buf []byte, part1 bool) (int, error) {
	var floor int
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
	return floor, nil
}

func Day01Branchless(buf []byte, part1 bool) (int, error) {
	var floor int
	for i, b := range buf {
		floor += 81 - 2*int(b)
		if !part1 && floor < 0 {
			// position is 1-based
			return i + 1, nil
		}
	}
	return floor, nil
}
