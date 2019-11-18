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
