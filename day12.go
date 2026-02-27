package adventofcode2015

import (
	"encoding/json"
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

func sum2(s string) int {
	buf := []byte(s)
	// "the empty interface says nothing"
	var data interface{}
	if err := json.Unmarshal(buf, &data); err != nil {
		return 0
	}

	numbers := make(chan int)
	go walk(data, numbers)

	var total int
	for n := range numbers {
		total += n
	}
	return total
}

func walk(root interface{}, numbers chan<- int) {
	defer close(numbers)
	stack := []interface{}{root}
	for len(stack) > 0 {
		last := len(stack) - 1
		v := stack[last]
		stack = stack[:last]

		switch vv := v.(type) {
		case string:
			// ignore strings
		case float64:
			numbers <- int(vv)
		case []interface{}:
			for i := len(vv) - 1; i >= 0; i-- {
				stack = append(stack, vv[i])
			}
		case map[string]interface{}:
			if hasRedProperty(vv) {
				continue
			}
			for _, item := range vv {
				stack = append(stack, item)
			}
		}
	}
}

func hasRedProperty(m map[string]interface{}) bool {
	for _, v := range m {
		switch t := v.(type) {
		case string:
			if t == "red" {
				return true
			}
		}
	}
	return false
}
