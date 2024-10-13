package adventofcode2015

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Day4Part1 returns lowest positive number whose MD5 starts with 5 zeroes.
func Day4Part1(s string) uint {
	return mine(s, 5)
}

// Day4Part2 returns lowest positive number whose MD5 starts with 6 zeroes.
func Day4Part2(s string) uint {
	return mine(s, 6)
}

func mine(s string, zeroes uint) (n uint) {
	prefix := strings.Repeat("0", int(zeroes))
	for !strings.HasPrefix(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, n)))), prefix) {
		n++
	}
	return
}
