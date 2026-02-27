package adventofcode2015

import "strconv"

const day17Storage = 150

type Day17Puzzle []uint

func NewDay17(lines []string) (Day17Puzzle, error) {
	ns := make(Day17Puzzle, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return nil, err
		}
		ns = append(ns, uint(n))
	}
	return ns, nil
}

// Day17 solves day 17 for the selected part.
func Day17(puzzle Day17Puzzle, part1 bool) uint {
	return day17Count(day17Storage, puzzle, part1)
}

func day17Count(storage uint, capacities []uint, part1 bool) uint {
	n := len(capacities)
	if n == 0 {
		return 0
	}

	var matches uint
	minContainers := n + 1
	minCount := uint(0)

	for mask := 0; mask < (1 << n); mask++ {
		sum := uint(0)
		count := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			sum += capacities[i]
			count++
		}
		if sum != storage {
			continue
		}
		if part1 {
			matches++
			continue
		}
		if count < minContainers {
			minContainers = count
			minCount = 1
		} else if count == minContainers {
			minCount++
		}
	}

	if part1 {
		return matches
	}
	return minCount
}

