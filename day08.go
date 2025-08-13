package adventofcode2015

// Day8Part1 returns logical number of characters for given buffer.
// Physical number of characters can easily be computed via len(buf).
func Day8Part1(buf []byte) int {
	n := 0
	for i := 0; i < len(buf); i++ {
		switch buf[i] {
		case '"':
		case '\\':
			i++
			switch buf[i] {
			case 'x':
				i += 2
				n++
			default:
				n++
			}
		default:
			n++
		}
	}
	return n
}

// Day8Part2 escapes a buffer and returns the number of escaped bytes, including
// a starting and leading ".
// " -> \"
// \ -> \\
func Day8Part2(buf []byte) int {
	total := len(buf)
	for _, b := range buf {
		if b == '\\' {
			total++
		}
		if b == '"' {
			total++
		}
	}
	// leading and trailing "
	return total + 2
}
