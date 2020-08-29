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
	go walk(0, data, numbers)

	var total int
	for n := range numbers {
		total += n
	}
	return total
}

func walk(depth int, v interface{}, numbers chan<- int) {
	switch vv := v.(type) {
	case string:
		// ignore strings
	case float64:
		// fmt.Println(v, "is number", vv)
		numbers <- int(vv)
	case []interface{}:
		// fmt.Println(v, "is an array:")
		for _, v := range vv {
			walk(depth+1, v, numbers)
		}
	case map[string]interface{}:
		// fmt.Println(v, "is a map:")
		if !hasRedProperty(vv) {
			for _, v := range vv {
				walk(depth+1, v, numbers)
			}
		}
	default:
		panic(vv)
	}
	// Once we're done, and back to root level, indicate end of numbers
	if depth == 0 {
		close(numbers)
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
