package adventofcode2015

func day03Key(x, y int) uint64 {
	return uint64(uint32(x))<<32 | uint64(uint32(y))
}

// Day03 solves day 3 for the selected part.
func Day03(buf []byte, part1 bool) (uint, error) {
	houses := make(map[uint64]struct{}, len(buf)+1)
	if part1 {
		x, y := 0, 0
		houses[day03Key(x, y)] = struct{}{}
		for _, b := range buf {
			switch b {
			case '>':
				x++
			case '<':
				x--
			case 'v':
				y++
			case '^':
				y--
			}
			houses[day03Key(x, y)] = struct{}{}
		}
		return uint(len(houses)), nil
	}

	sx, sy := 0, 0
	rx, ry := 0, 0
	who := 0
	houses[day03Key(0, 0)] = struct{}{}
	for _, b := range buf {
		if who == 0 {
			switch b {
			case '>':
				sx++
			case '<':
				sx--
			case 'v':
				sy++
			case '^':
				sy--
			}
			houses[day03Key(sx, sy)] = struct{}{}
		} else {
			switch b {
			case '>':
				rx++
			case '<':
				rx--
			case 'v':
				ry++
			case '^':
				ry--
			}
			houses[day03Key(rx, ry)] = struct{}{}
		}
		who = 1 - who
	}
	return uint(len(houses)), nil
}
