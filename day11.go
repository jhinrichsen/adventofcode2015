package adventofcode2015

import "errors"

type Day11Puzzle string

func NewDay11(lines []string) (Day11Puzzle, error) {
	if len(lines) != 1 {
		return "", errors.New("invalid input")
	}
	return Day11Puzzle(lines[0]), nil
}

func Day11(puzzle Day11Puzzle, part1 bool) string {
	first := next(string(puzzle))
	if part1 {
		return first
	}
	return next(first)
}

func incBytes(bs []byte) {
	for idx := len(bs) - 1; idx >= 0; idx-- {
		if bs[idx] == 'z' {
			bs[idx] = 'a'
			continue
		}
		bs[idx]++
		return
	}
}

// Passwords must include one increasing straight of at least three letters.
func req1Bytes(bs []byte) bool {
	for i := 0; i+2 < len(bs); i++ {
		if bs[i]+1 == bs[i+1] && bs[i+1]+1 == bs[i+2] {
			return true
		}
	}
	return false
}

// Passwords may not contain the letters i, o, or l.
func req2Bytes(bs []byte) bool {
	for i := range len(bs) {
		if bs[i] == 'i' || bs[i] == 'o' || bs[i] == 'l' {
			return false
		}
	}
	return true
}

// Passwords must contain at least two different, non-overlapping pairs of
// letters.
func req3Bytes(bs []byte) bool {
	var pairs [26]bool
	nPairs := 0
	for i := 0; i+1 < len(bs); i++ {
		if bs[i] == bs[i+1] {
			idx := bs[i] - 'a'
			if !pairs[idx] {
				pairs[idx] = true
				nPairs++
				if nPairs >= 2 {
					return true
				}
			}
			// pairs must not overlap
			i++
		}
	}
	return false
}

func next(s string) string {
	bs := []byte(s)
	for {
		incBytes(bs)
		if req1Bytes(bs) && req2Bytes(bs) && req3Bytes(bs) {
			return string(bs)
		}
	}
}
