package adventofcode2015

func inc(s string) string {
	bs := []byte(s)
	idx := len(bs) - 1
_inc:
	if bs[idx] == 'z' {
		bs[idx] = 'a'
		idx--
		goto _inc
	} else {
		bs[idx]++
	}
	return string(bs)
}

// Passwords must include one increasing straight of at least three letters.
func req1(s string) bool {
	bs := []byte(s)
	hasThree := func(idx int) bool {
		return bs[idx] == bs[idx+1]-1 &&
			bs[idx+1] == bs[idx+2]-1
	}
	for i := range bs[:len(bs)-2] {
		if hasThree(i) {
			return true
		}
	}
	return false
}

// Passwords may not contain the letters i, o, or l.
func req2(s string) bool {
	for _, r := range s {
		if r == 'i' || r == 'o' || r == 'l' {
			return false
		}
	}
	return true
}

// Passwords must contain at least two different, non-overlapping pairs of
// letters.
func req3(s string) bool {
	pairs := make(map[byte]bool)
	bs := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if bs[i] == bs[i+1] {
			pairs[bs[i]] = true
			// pars must not overlap
			i++
		}
	}
	return len(pairs) >= 2
}

func next(s string) string {
_again:
	s = inc(s)
	if !req1(s) || !req2(s) || !req3(s) {
		goto _again
	}
	return s
}
