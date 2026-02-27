package adventofcode2015

import "image"

func move(pos image.Point, b byte) image.Point {
	switch b {
	case '>':
		pos.X++
	case '<':
		pos.X--
	case 'v':
		pos.Y++
	case '^':
		pos.Y--
	}
	return pos
}

// Day03 solves day 3 for the selected part.
func Day03(buf []byte, part1 bool) (uint, error) {
	houses := make(map[image.Point]bool)
	if part1 {
		pos := image.Point{X: 0, Y: 0}
		houses[pos] = true
		for _, b := range buf {
			pos = move(pos, b)
			houses[pos] = true
		}
		return uint(len(houses)), nil
	}

	// index santa = 0, robo santa = 1
	poss := []image.Point{{X: 0, Y: 0}, {X: 0, Y: 0}}
	who := 0
	houses[poss[0]] = true
	for _, b := range buf {
		poss[who] = move(poss[who], b)
		houses[poss[who]] = true
		who = 1 - who
	}
	return uint(len(houses)), nil
}
