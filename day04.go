package adventofcode2015

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Day04 solves day 4 for the selected part.
func Day04(buf []byte, part1 bool) (uint, error) {
	secret := strings.TrimSpace(string(buf))
	if part1 {
		return mine(secret, 5), nil
	}
	return mine(secret, 6), nil
}

func mine(s string, zeroes uint) (n uint) {
	prefix := strings.Repeat("0", int(zeroes))
	for !strings.HasPrefix(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, n)))), prefix) {
		n++
	}
	return
}
