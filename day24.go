package adventofcode2015

import (
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

func quantumEntanglement(a []uint) uint {
	qe := uint(1)
	for _, w := range a {
		qe *= w
	}
	return qe
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
	total := ws.sum()
	if total%nGroups != 0 {
		return 0
	}

	ch := make(chan []uint)
	go heapUint(len(ws), ws, ch)

	groupWeight := total / nGroups
	minPackets := len(ws)
	minQe := quantumEntanglement(ws)
	for prospect := range ch {
		groupWeights := make([]uint, nGroups)
		j := 0
		nPacketsGroup1 := 0
		for i := 0; i < len(prospect); i++ {
			groupWeights[j] += prospect[i]
			if groupWeights[j] > groupWeight {
				break
			}
			// group 1 complete, matching, entering group 2
			if groupWeights[j] == groupWeight {
				if j == 0 {
					nPacketsGroup1 = i + 1
				}
				j++
				// distributed into 3 equal groups?
				if j+1 == len(groupWeights) {
					// level 1: number of packets
					if nPacketsGroup1 < minPackets {
						minPackets = nPacketsGroup1
						// new local min for this number of packets
						minQe = quantumEntanglement(prospect[0:nPacketsGroup1])
						continue
					}
					// level 2: qe
					if nPacketsGroup1 == minPackets {
						qe := quantumEntanglement(prospect[0:nPacketsGroup1])
						if qe < minQe {
							minQe = qe
						}
					}
				}
			}
		}

	}
	return minQe
}
