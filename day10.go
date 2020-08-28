package adventofcode2015

import (
	"strconv"
	"strings"
)

func lookAndSay(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return "1" + s
	}
	var sb strings.Builder
	bs := []byte(s)
	l := len(bs)
	same := func(idx int) int {
		j := idx + 1
		for ; j < len(bs) && bs[idx] == bs[j]; j++ {
		}
		return j - idx
	}
	for i := 0; i < l; {
		n := same(i)
		// number of same digits
		sb.WriteString(strconv.Itoa(n))
		// digit itself
		sb.WriteByte(bs[i])

		i += n
	}
	return sb.String()
}
