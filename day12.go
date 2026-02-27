package adventofcode2015

import "io"

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
	type frame struct {
		delim        byte
		sum          int
		skipObject   bool
		expectingKey bool
	}
	stack := make([]frame, 0, 64)
	total := 0
	addValue := func(v int) {
		if len(stack) == 0 {
			total += v
			return
		}
		stack[len(stack)-1].sum += v
	}
	setObjectValueDone := func() {
		if len(stack) == 0 {
			return
		}
		last := len(stack) - 1
		if stack[last].delim == '{' && !stack[last].expectingKey {
			stack[last].expectingKey = true
		}
	}
	for i := 0; i < len(buf); {
		switch buf[i] {
		case '{':
			stack = append(stack, frame{delim: '{', expectingKey: true})
			i++
		case '[':
			stack = append(stack, frame{delim: '['})
			i++
		case '}', ']':
			last := len(stack) - 1
			value := 0
			if !(stack[last].delim == '{' && stack[last].skipObject) {
				value = stack[last].sum
			}
			stack = stack[:last]
			addValue(value)
			setObjectValueDone()
			i++
		case '"':
			i++
			start := i
			for i < len(buf) {
				if buf[i] == '\\' {
					i += 2
					continue
				}
				if buf[i] == '"' {
					break
				}
				i++
			}
			if i >= len(buf) {
				return 0, io.ErrUnexpectedEOF
			}
			if len(stack) == 0 || stack[len(stack)-1].delim != '{' {
				i++
				continue
			}
			last := len(stack) - 1
			if stack[last].expectingKey {
				stack[last].expectingKey = false
				i++
				continue
			}
			if i-start == 3 && buf[start] == 'r' && buf[start+1] == 'e' && buf[start+2] == 'd' {
				stack[last].skipObject = true
			}
			stack[last].expectingKey = true
			i++
		case 't':
			i += 4
			setObjectValueDone()
		case 'f':
			i += 5
			setObjectValueDone()
		case 'n':
			i += 4
			setObjectValueDone()
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sign := 1
			if buf[i] == '-' {
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
			addValue(sign * n)
			setObjectValueDone()
			i++
		default:
			i++
		}
	}
	return total, nil
}
