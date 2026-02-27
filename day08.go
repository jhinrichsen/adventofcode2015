package adventofcode2015

// Day08 solves day 8 for the selected part.
func Day08(lines []string, part1 bool) uint {
	var totalCode uint
	var totalMemory uint
	var totalEncoded uint

	for _, line := range lines {
		codeLen := uint(len(line))
		totalCode += codeLen
		if part1 {
			totalMemory += day08MemoryLen(line)
			continue
		}
		totalEncoded += day08EncodedLen(line)
	}

	if part1 {
		return totalCode - totalMemory
	}
	return totalEncoded - totalCode
}

func day08MemoryLen(s string) uint {
	if len(s) < 2 {
		return 0
	}

	var n uint
	for i := 1; i < len(s)-1; i++ {
		if s[i] != '\\' {
			n++
			continue
		}
		i++
		if i >= len(s)-1 {
			break
		}
		if s[i] == 'x' {
			i += 2
		}
		n++
	}
	return n
}

func day08EncodedLen(s string) uint {
	n := uint(2) // opening + closing quotes
	for i := range len(s) {
		if s[i] == '\\' || s[i] == '"' {
			n += 2
			continue
		}
		n++
	}
	return n
}
