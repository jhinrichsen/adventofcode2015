package adventofcode2015

import (
	"math/bits"
	"slices"
	"strconv"
)

type weights []uint

type Day24Puzzle weights

func NewDay24(lines []string) (Day24Puzzle, error) {
	ws, err := newWeights(lines)
	if err != nil {
		return nil, err
	}
	return Day24Puzzle(ws), nil
}

func (a weights) sum() (n uint) {
	for i := range a {
		n += a[i]
	}
	return
}

func newWeights(lines []string) (weights, error) {
	var ws weights
	for _, line := range lines {
		x, err := strconv.Atoi(line)
		if err != nil {
			return ws, err
		}
		// weights can never be negative (helium balloons, anyone?)
		ws = append(ws, uint(x))
	}
	return ws, nil
}

// Day24 solves day 24 for the selected part.
func Day24(puzzle Day24Puzzle, part1 bool) uint {
	if part1 {
		return day24(weights(puzzle), 3)
	}
	return day24(weights(puzzle), 4)
}

func day24(ws weights, nGroups uint) uint {
	if len(ws) == 0 || len(ws) > 63 {
		return 0
	}

	total := ws.sum()
	if total%nGroups != 0 {
		return 0
	}

	groupWeight := total / nGroups

	ordered := slices.Clone([]uint(ws))
	slices.Sort(ordered)
	slices.Reverse(ordered)

	targetSubsets := day24TargetSubsets(ordered, groupWeight)
	if len(targetSubsets) == 0 {
		return 0
	}

	candidates := make([]day24Candidate, 0, len(targetSubsets))
	for _, mask := range targetSubsets {
		candidates = append(candidates, day24Candidate{
			mask:  mask,
			count: bits.OnesCount64(mask),
			qe:    day24SubsetQE(mask, ordered),
		})
	}

	slices.SortFunc(candidates, func(a, b day24Candidate) int {
		if a.count != b.count {
			if a.count < b.count {
				return -1
			}
			return 1
		}
		if a.qe != b.qe {
			if a.qe < b.qe {
				return -1
			}
			return 1
		}
		if a.mask < b.mask {
			return -1
		}
		if a.mask > b.mask {
			return 1
		}
		return 0
	})

	allMask := (uint64(1) << len(ordered)) - 1
	for _, candidate := range candidates {
		remaining := allMask ^ candidate.mask
		if day24CanPartition(remaining, nGroups-1, targetSubsets) {
			return candidate.qe
		}
	}

	return 0
}

type day24Candidate struct {
	mask  uint64
	count int
	qe    uint
}

type day24Node struct {
	idx  int
	sum  uint
	mask uint64
}

func day24TargetSubsets(ws []uint, target uint) []uint64 {
	n := len(ws)
	suffix := make([]uint, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + ws[i]
	}

	stack := []day24Node{{}}
	subsets := make([]uint64, 0, 4096)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.sum == target {
			subsets = append(subsets, cur.mask)
			continue
		}
		if cur.idx == n || cur.sum > target {
			continue
		}
		if cur.sum+suffix[cur.idx] < target {
			continue
		}

		nextIdx := cur.idx + 1
		stack = append(stack, day24Node{idx: nextIdx, sum: cur.sum, mask: cur.mask})
		with := cur.sum + ws[cur.idx]
		if with <= target {
			stack = append(stack, day24Node{
				idx:  nextIdx,
				sum:  with,
				mask: cur.mask | (uint64(1) << cur.idx),
			})
		}
	}

	return subsets
}

func day24SubsetQE(mask uint64, ws []uint) uint {
	qe := uint(1)
	for i, w := range ws {
		if mask&(uint64(1)<<i) != 0 {
			qe *= w
		}
	}
	return qe
}

func day24CanPartition(mask uint64, groups uint, targetSubsets []uint64) bool {
	switch groups {
	case 1:
		return true
	case 2:
		return day24HasTargetSubset(mask, targetSubsets)
	case 3:
		anchor := mask & -mask
		for _, subset := range targetSubsets {
			if subset&anchor == 0 {
				continue
			}
			if subset&mask != subset {
				continue
			}
			remaining := mask ^ subset
			if day24HasTargetSubset(remaining, targetSubsets) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func day24HasTargetSubset(mask uint64, targetSubsets []uint64) bool {
	for _, subset := range targetSubsets {
		if subset&mask == subset {
			return true
		}
	}
	return false
}
