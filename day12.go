package adventofcode2015

import (
	"encoding/json"
	"strconv"
)

// Day12 solves day 12 for the selected part.
func Day12(buf []byte, part1 bool) (uint, error) {
	if part1 {
		return uint(sumNumbers(buf)), nil
	}
	n, err := sumWithoutRed(buf)
	return uint(n), err
}

func sumNumbers(buf []byte) int {
	total := 0
	for i := 0; i < len(buf); i++ {
		b := buf[i]
		if (b < '0' || b > '9') && b != '-' {
			continue
		}
		sign := 1
		if b == '-' {
			sign = -1
			i++
			if i >= len(buf) || buf[i] < '0' || buf[i] > '9' {
				continue
			}
		}
		n := int(buf[i] - '0')
		for i+1 < len(buf) && buf[i+1] >= '0' && buf[i+1] <= '9' {
			i++
			n = n*10 + int(buf[i]-'0')
		}
		total += sign * n
	}
	return total
}

func sumWithoutRed(buf []byte) (int, error) {
	var root any
	if err := json.Unmarshal(buf, &root); err != nil {
		return 0, err
	}

	total := 0
	stack := []any{root}
	for len(stack) > 0 {
		last := len(stack) - 1
		v := stack[last]
		stack = stack[:last]

		switch t := v.(type) {
		case float64:
			total += int(t)
		case []any:
			for i := len(t) - 1; i >= 0; i-- {
				stack = append(stack, t[i])
			}
		case map[string]any:
			if hasRedValue(t) {
				continue
			}
			for _, item := range t {
				stack = append(stack, item)
			}
		case json.Number:
			n, err := strconv.Atoi(string(t))
			if err == nil {
				total += n
			}
		}
	}
	return total, nil
}

func hasRedValue(m map[string]any) bool {
	for _, v := range m {
		s, ok := v.(string)
		if ok && s == "red" {
			return true
		}
	}
	return false
}
