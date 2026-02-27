package adventofcode2015

import (
	"bytes"
	"crypto/md5"
	"strconv"
)

// Day04 solves day 4 for the selected part.
func Day04(buf []byte, part1 bool) (uint, error) {
	secret := bytes.TrimSpace(buf)
	if part1 {
		return mine(secret, true), nil
	}
	return mine(secret, false), nil
}

func mine(secret []byte, part1 bool) (n uint) {
	candidate := make([]byte, len(secret), len(secret)+20)
	copy(candidate, secret)
	for {
		candidate = candidate[:len(secret)]
		candidate = strconv.AppendUint(candidate, uint64(n), 10)
		sum := md5.Sum(candidate)
		if part1 {
			if sum[0] == 0 && sum[1] == 0 && sum[2] < 0x10 {
				return
			}
		} else if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			return
		}
		n++
	}
}
