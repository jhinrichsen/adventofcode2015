package adventofcode2015

type position [2]int

// Day03Part1 returns number of fields visited.
func Day03Part1(buf []byte) uint {
	houses := make(map[position]bool)
	pos := position{0, 0}
	houses[pos] = true
	for _, b := range buf {
		pos = move(pos, b)
		houses[pos] = true
	}
	return uint(len(houses))
}

func move(pos position, b byte) position {
	switch b {
	case '>':
		pos[0]++
	case '<':
		pos[0]--
	case 'v':
		pos[1]++
	case '^':
		pos[1]--
	}
	return pos
}

// Day03Part2 TODO.
func Day03Part2(buf []byte) uint {
	houses := make(map[position]bool)
	// index santa = 0, robo santa = 1
	poss := []position{
		{0, 0},
		{0, 0},
	}
	// santa starts
	who := 0
	mark := func() {
		houses[poss[who]] = true
	}
	// deliver to current position
	mark()
	for _, b := range buf {
		poss[who] = move(poss[who], b)
		mark()
		// take turns
		who = 1 - who
	}
	return uint(len(houses))
}
