package adventofcode2015

import (
	"fmt"
	"strings"
)

const day07MaxWires = 26 + 26*26 // a..z, aa..zz

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
	id      int
	literal uint16
	isValue bool
}

type day07Expr struct {
	op day07Op
	a  day07Operand
	b  day07Operand
}

type Day07Puzzle struct {
	exprs   [day07MaxWires]day07Expr
	defined [day07MaxWires]bool
	aID     int
	bID     int
}

func day07WireID(s string) (int, error) {
	if len(s) == 1 {
		c := s[0]
		if c < 'a' || c > 'z' {
			return -1, fmt.Errorf("invalid wire name %q", s)
		}
		return int(c - 'a'), nil
	}
	if len(s) == 2 {
		c0, c1 := s[0], s[1]
		if c0 < 'a' || c0 > 'z' || c1 < 'a' || c1 > 'z' {
			return -1, fmt.Errorf("invalid wire name %q", s)
		}
		return 26 + int(c0-'a')*26 + int(c1-'a'), nil
	}
	return -1, fmt.Errorf("invalid wire name %q", s)
}

func parseDay07Literal(s string) (uint16, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty literal")
	}
	var n uint32
	for i := range len(s) {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid literal %q", s)
		}
		n = n*10 + uint32(c-'0')
		if n > 65535 {
			return 0, fmt.Errorf("literal out of range %q", s)
		}
	}
	return uint16(n), nil
}

func parseDay07Operand(s string) (day07Operand, error) {
	if len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		v, err := parseDay07Literal(s)
		if err != nil {
			return day07Operand{}, err
		}
		return day07Operand{literal: v, isValue: true}, nil
	}
	id, err := day07WireID(s)
	if err != nil {
		return day07Operand{}, err
	}
	return day07Operand{id: id}, nil
}

func parseDay07Line(line string) (destID int, expr day07Expr, err error) {
	left, dest, ok := strings.Cut(line, " -> ")
	if !ok {
		return -1, day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
	}
	destID, err = day07WireID(dest)
	if err != nil {
		return -1, day07Expr{}, err
	}

	if strings.HasPrefix(left, "NOT ") {
		a, err := parseDay07Operand(left[4:])
		if err != nil {
			return -1, day07Expr{}, err
		}
		return destID, day07Expr{op: day07Not, a: a}, nil
	}

	if aS, bS, ok := strings.Cut(left, " AND "); ok {
		a, err := parseDay07Operand(aS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		b, err := parseDay07Operand(bS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		return destID, day07Expr{op: day07And, a: a, b: b}, nil
	}
	if aS, bS, ok := strings.Cut(left, " OR "); ok {
		a, err := parseDay07Operand(aS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		b, err := parseDay07Operand(bS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		return destID, day07Expr{op: day07Or, a: a, b: b}, nil
	}
	if aS, bS, ok := strings.Cut(left, " LSHIFT "); ok {
		a, err := parseDay07Operand(aS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		b, err := parseDay07Operand(bS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		return destID, day07Expr{op: day07LShift, a: a, b: b}, nil
	}
	if aS, bS, ok := strings.Cut(left, " RSHIFT "); ok {
		a, err := parseDay07Operand(aS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		b, err := parseDay07Operand(bS)
		if err != nil {
			return -1, day07Expr{}, err
		}
		return destID, day07Expr{op: day07RShift, a: a, b: b}, nil
	}

	a, err := parseDay07Operand(left)
	if err != nil {
		return -1, day07Expr{}, err
	}
	return destID, day07Expr{op: day07Assign, a: a}, nil
}

func NewDay07(lines []string) (Day07Puzzle, error) {
	puzzle := Day07Puzzle{aID: -1, bID: -1}
	for _, line := range lines {
		destID, expr, err := parseDay07Line(line)
		if err != nil {
			return Day07Puzzle{}, err
		}
		puzzle.exprs[destID] = expr
		puzzle.defined[destID] = true
	}
	aID, err := day07WireID("a")
	if err != nil {
		return Day07Puzzle{}, err
	}
	bID, err := day07WireID("b")
	if err != nil {
		return Day07Puzzle{}, err
	}
	if puzzle.defined[aID] {
		puzzle.aID = aID
	}
	if puzzle.defined[bID] {
		puzzle.bID = bID
	}
	return puzzle, nil
}

func day07DepIDs(expr day07Expr) [2]int {
	ids := [2]int{-1, -1}
	switch expr.op {
	case day07Assign, day07Not:
		if !expr.a.isValue {
			ids[0] = expr.a.id
		}
	case day07And, day07Or, day07LShift, day07RShift:
		if !expr.a.isValue {
			ids[0] = expr.a.id
		}
		if !expr.b.isValue {
			ids[1] = expr.b.id
		}
	}
	return ids
}

func day07EvalExpr(expr day07Expr, vals *[day07MaxWires]uint16) uint16 {
	operand := func(op day07Operand) uint16 {
		if op.isValue {
			return op.literal
		}
		return vals[op.id]
	}
	x := operand(expr.a)
	switch expr.op {
	case day07Assign:
		return x
	case day07Not:
		return ^x
	case day07And:
		return x & operand(expr.b)
	case day07Or:
		return x | operand(expr.b)
	case day07LShift:
		return x << operand(expr.b)
	case day07RShift:
		return x >> operand(expr.b)
	default:
		return 0
	}
}

func (a Day07Puzzle) signal(wire string, overrideB *uint16) (uint16, error) {
	targetID, err := day07WireID(wire)
	if err != nil {
		return 0, err
	}
	if !a.defined[targetID] {
		return 0, fmt.Errorf("unknown wire %q", wire)
	}

	var state [day07MaxWires]byte // 0=new, 1=visiting, 2=done
	var vals [day07MaxWires]uint16
	var stack [day07MaxWires]int
	sp := 0
	stack[sp] = targetID
	sp++

	for sp > 0 {
		id := stack[sp-1]
		if state[id] == 2 {
			sp--
			continue
		}
		if overrideB != nil && id == a.bID {
			vals[id] = *overrideB
			state[id] = 2
			sp--
			continue
		}
		if !a.defined[id] {
			return 0, fmt.Errorf("unknown wire %q", wire)
		}
		if state[id] == 0 {
			state[id] = 1
		}

		expr := a.exprs[id]
		deps := day07DepIDs(expr)
		needDep := false
		for _, depID := range deps {
			if depID < 0 {
				continue
			}
			switch state[depID] {
			case 0:
				stack[sp] = depID
				sp++
				needDep = true
			case 1:
				return 0, fmt.Errorf("circular dependency for wire %q", wire)
			}
		}
		if needDep {
			continue
		}

		vals[id] = day07EvalExpr(expr, &vals)
		state[id] = 2
		sp--
	}
	return vals[targetID], nil
}

// Day07Part1 returns the signal provided to wire "a".
func Day07Part1(lines []string) (uint16, error) {
	puzzle, err := NewDay07(lines)
	if err != nil {
		return 0, err
	}
	if puzzle.aID < 0 {
		return 0, fmt.Errorf("wire %q not found", "a")
	}
	return puzzle.signal("a", nil)
}

// Day07Part2 overrides wire "b" with the Part 1 result and recomputes wire "a".
func Day07Part2(lines []string) (uint16, error) {
	puzzle, err := NewDay07(lines)
	if err != nil {
		return 0, err
	}
	if puzzle.aID < 0 {
		return 0, fmt.Errorf("wire %q not found", "a")
	}
	if puzzle.bID < 0 {
		return 0, fmt.Errorf("wire %q not found", "b")
	}
	aSignal, err := puzzle.signal("a", nil)
	if err != nil {
		return 0, err
	}
	return puzzle.signal("a", &aSignal)
}
