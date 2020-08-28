package adventofcode2015

import (
	"strconv"
	"strings"
	"unicode"
)

func isNumeric(r rune) bool {
	return r == '+' ||
		r == '-' ||
		unicode.IsDigit(r)
}

func sum(s string) int {
	// erase everything non-numeric
	var sb strings.Builder
	for _, r := range s {
		if isNumeric(r) {
			sb.WriteRune(r)
		} else {
			sb.WriteByte(' ')
		}
	}

	// Iterate numbers
	var total int
	for _, s := range strings.Fields(sb.String()) {
		n, _ := strconv.Atoi(s)
		total += n
	}
	return total
}
