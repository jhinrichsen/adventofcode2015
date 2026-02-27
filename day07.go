package adventofcode2015

import (
	"fmt"
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
	wireID map[string]int
	exprs  []day07Expr
	aID    int
	bID    int
}

func parseDay07Operand(s string, getID func(string) int) day07Operand {
	v, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		return day07Operand{literal: uint16(v), isValue: true}
	}
	return day07Operand{id: getID(s)}
}

func parseDay07Line(line string,
	getID func(string) int) (destID int, expr day07Expr, err error) {
	parts := strings.Fields(line)
	switch len(parts) {
	case 3:
		if parts[1] != "->" {
			return -1, day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
		}
		return getID(parts[2]), day07Expr{
			op: day07Assign,
			a:  parseDay07Operand(parts[0], getID),
		}, nil
	case 4:
		if parts[0] != "NOT" || parts[2] != "->" {
			return -1, day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
		}
		return getID(parts[3]), day07Expr{
			op: day07Not,
			a:  parseDay07Operand(parts[1], getID),
		}, nil
	case 5:
		if parts[3] != "->" {
			return -1, day07Expr{}, fmt.Errorf("invalid instruction: %q", line)
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
			return -1, day07Expr{}, fmt.Errorf("unknown operator %q", parts[1])
		}
		return getID(parts[4]), day07Expr{
			op: op,
			a:  parseDay07Operand(parts[0], getID),
			b:  parseDay07Operand(parts[2], getID),
		}, nil
	default:
		return -1, day07Expr{}, fmt.Errorf("cannot parse instruction: %q", line)
	}
}

func NewDay07(lines []string) (Day07Puzzle, error) {
	puzzle := Day07Puzzle{
		wireID: make(map[string]int, len(lines)),
		exprs:  make([]day07Expr, 0, len(lines)),
		aID:    -1,
		bID:    -1,
	}
	getID := func(name string) int {
		if id, ok := puzzle.wireID[name]; ok {
			return id
		}
		id := len(puzzle.exprs)
		puzzle.wireID[name] = id
		puzzle.exprs = append(puzzle.exprs, day07Expr{})
		return id
	}

	for _, line := range lines {
		destID, expr, err := parseDay07Line(line, getID)
		if err != nil {
			return Day07Puzzle{}, err
		}
		puzzle.exprs[destID] = expr
	}
	if id, ok := puzzle.wireID["a"]; ok {
		puzzle.aID = id
	}
	if id, ok := puzzle.wireID["b"]; ok {
		puzzle.bID = id
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

func day07EvalExpr(expr day07Expr, vals []uint16) uint16 {
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
	targetID, ok := a.wireID[wire]
	if !ok {
		return 0, fmt.Errorf("unknown wire %q", wire)
	}

	state := make([]byte, len(a.exprs)) // 0=new, 1=visiting, 2=done
	vals := make([]uint16, len(a.exprs))
	stack := make([]int, 0, 64)
	stack = append(stack, targetID)

	for len(stack) > 0 {
		id := stack[len(stack)-1]
		if state[id] == 2 {
			stack = stack[:len(stack)-1]
			continue
		}
		if overrideB != nil && id == a.bID {
			vals[id] = *overrideB
			state[id] = 2
			stack = stack[:len(stack)-1]
			continue
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
				stack = append(stack, depID)
				needDep = true
			case 1:
				return 0, fmt.Errorf("circular dependency for wire %q", wire)
			}
		}
		if needDep {
			continue
		}

		vals[id] = day07EvalExpr(expr, vals)
		state[id] = 2
		stack = stack[:len(stack)-1]
	}
	return vals[targetID], nil
}

func (a Day07Puzzle) Signal(wire string) (uint16, error) {
	return a.signal(wire, nil)
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
