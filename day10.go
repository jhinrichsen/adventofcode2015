package adventofcode2015

import (
	"errors"
	"strconv"
	"strings"
)

type Day10Puzzle string

func NewDay10(lines []string) (Day10Puzzle, error) {
	if len(lines) != 1 {
		return "", errors.New("invalid input")
	}
	return Day10Puzzle(lines[0]), nil
}

func Day10(puzzle Day10Puzzle, part1 bool) uint {
	s := string(puzzle)
	steps := 50
	if part1 {
		steps = 40
	}
	for range steps {
		s = lookAndSay(s)
	}
	return uint(len(s))
}

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
