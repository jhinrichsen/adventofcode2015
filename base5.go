package adventofcode2015

// Base5 digits.
type Base5 struct {
	// order is right to left so we can easily add new digits to the right.
	Buf []byte
}

// NewBase5 creates 0 in base 5 format.
func NewBase5(digits uint) Base5 {
	b := Base5{}
	b.Buf = make([]byte, digits)
	return b
}

// Inc increments a base 5 digit by one.
func (a *Base5) Inc() {
	i := 0
inc:
	if a.Buf[i] == 4 {
		a.Buf[i] = 0
		if i == len(a.Buf) {
			a.Buf = append(a.Buf, 0)
		}
		i++
		goto inc
	} else {
		a.Buf[i]++
	}
}
