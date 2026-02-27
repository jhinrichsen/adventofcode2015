package adventofcode2015

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type day07Op uint8

const (
	day07Assign day07Op = iota
	day07Not
	day07And
	day07Or
	day07LShift
	day07RShift
)

type day07Operand struct {
	wire    string
	literal uint16
	isValue bool
}

type day07Expr struct {
	op day07Op
	a  day07Operand
	b  day07Operand
}

type Day07Puzzle struct {
	wires map[string]day07Expr
}

type day07Circuit struct {
	wires map[string]day07Expr
	cache map[string]uint16
}

func parseDay07Operand(s string) day07Operand {
	v, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		return day07Operand{literal: uint16(v), isValue: true}
	}
	return day07Operand{wire: s}
}

func parseDay07Line(line string) (dest string, expr day07Expr, err error) {
	parts := strings.Fields(line)
	switch len(parts) {
	case 3:
		if parts[1] != "->" {
			return "", day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
		}
		return parts[2], day07Expr{
			op: day07Assign,
			a:  parseDay07Operand(parts[0]),
		}, nil
	case 4:
		if parts[0] != "NOT" || parts[2] != "->" {
			return "", day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
		}
		return parts[3], day07Expr{
			op: day07Not,
			a:  parseDay07Operand(parts[1]),
		}, nil
	case 5:
		if parts[3] != "->" {
			return "", day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
		}
		var op day07Op
		switch parts[1] {
		case "AND":
			op = day07And
		case "OR":
			op = day07Or
		case "LSHIFT":
			op = day07LShift
		case "RSHIFT":
			op = day07RShift
		default:
			return "", day07Expr{}, fmt.Errorf("unknown operator %q", parts[1])
		}
		return parts[4], day07Expr{
			op: op,
			a:  parseDay07Operand(parts[0]),
			b:  parseDay07Operand(parts[2]),
		}, nil
	default:
		return "", day07Expr{}, fmt.Errorf("cannot parse instruction: %q", line)
	}
}

func NewDay07(lines []string) (Day07Puzzle, error) {
	puzzle := Day07Puzzle{
		wires: make(map[string]day07Expr, len(lines)),
	}
	for _, line := range lines {
		dest, expr, err := parseDay07Line(line)
		if err != nil {
			return Day07Puzzle{}, err
		}
		puzzle.wires[dest] = expr
	}
	return puzzle, nil
}

func day07CircuitFromPuzzle(puzzle Day07Puzzle) day07Circuit {
	return day07Circuit{
		wires: maps.Clone(puzzle.wires),
		cache: make(map[string]uint16, len(puzzle.wires)),
	}
}

func (a *day07Circuit) operand(op day07Operand) (uint16, error) {
	if op.isValue {
		return op.literal, nil
	}
	return a.signal(op.wire)
}

func (a *day07Circuit) signal(wire string) (uint16, error) {
	if v, ok := a.cache[wire]; ok {
		return v, nil
	}
	expr, ok := a.wires[wire]
	if !ok {
		return 0, fmt.Errorf("unknown wire %q", wire)
	}

	x, err := a.operand(expr.a)
	if err != nil {
		return 0, err
	}

	var v uint16
	switch expr.op {
	case day07Assign:
		v = x
	case day07Not:
		v = ^x
	case day07And:
		y, err := a.operand(expr.b)
		if err != nil {
			return 0, err
		}
		v = x & y
	case day07Or:
		y, err := a.operand(expr.b)
		if err != nil {
			return 0, err
		}
		v = x | y
	case day07LShift:
		y, err := a.operand(expr.b)
		if err != nil {
			return 0, err
		}
		v = x << y
	case day07RShift:
		y, err := a.operand(expr.b)
		if err != nil {
			return 0, err
		}
		v = x >> y
	default:
		return 0, fmt.Errorf("unknown op %d", expr.op)
	}

	a.cache[wire] = v
	return v, nil
}

// Day07Part1 returns the signal provided to wire "a".
func Day07Part1(lines []string) (uint16, error) {
	puzzle, err := NewDay07(lines)
	if err != nil {
		return 0, err
	}
	c := day07CircuitFromPuzzle(puzzle)
	return c.signal("a")
}

// Day07Part2 overrides wire "b" with the Part 1 result and recomputes wire "a".
func Day07Part2(lines []string) (uint16, error) {
	puzzle, err := NewDay07(lines)
	if err != nil {
		return 0, err
	}
	c1 := day07CircuitFromPuzzle(puzzle)
	a, err := c1.signal("a")
	if err != nil {
		return 0, err
	}
	c := day07CircuitFromPuzzle(puzzle)
	c.wires["b"] = day07Expr{
		op: day07Assign,
		a:  day07Operand{literal: a, isValue: true},
	}
	return c.signal("a")
}
