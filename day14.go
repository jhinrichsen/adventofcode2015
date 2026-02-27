package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

const day14RaceSeconds = 2503

type day14Reindeer struct {
	name     string
	velocity uint
	flying   uint
	resting  uint
}

type Day14Puzzle []day14Reindeer

func NewDay14(lines []string) (Day14Puzzle, error) {
	rs := make(Day14Puzzle, 0, len(lines))
	for i, line := range lines {
		r, err := day14ParseReindeer(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}
		rs = append(rs, r)
	}
	return rs, nil
}

// Day14 solves day 14 for the selected part.
func Day14(puzzle Day14Puzzle, part1 bool) uint {
	if part1 {
		return day14DistanceWinner(puzzle, day14RaceSeconds)
	}
	return day14ScoreWinner(puzzle, day14RaceSeconds)
}

func day14ParseReindeer(line string) (day14Reindeer, error) {
	fields := strings.Fields(line)
	if len(fields) != 15 {
		return day14Reindeer{}, fmt.Errorf("want 15 fields but got %d", len(fields))
	}

	v, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return day14Reindeer{}, err
	}
	f, err := strconv.ParseUint(fields[6], 10, 64)
	if err != nil {
		return day14Reindeer{}, err
	}
	r, err := strconv.ParseUint(fields[13], 10, 64)
	if err != nil {
		return day14Reindeer{}, err
	}

	return day14Reindeer{
		name:     fields[0],
		velocity: uint(v),
		flying:   uint(f),
		resting:  uint(r),
	}, nil
}

func day14Distance(r day14Reindeer, sec uint) uint {
	frame := r.flying + r.resting
	full := sec / frame
	partial := min(sec%frame, r.flying)
	return full*r.flying*r.velocity + partial*r.velocity
}

func day14DistanceWinner(rs []day14Reindeer, sec uint) uint {
	best := uint(0)
	for _, r := range rs {
		best = max(best, day14Distance(r, sec))
	}
	return best
}

func day14ScoreWinner(rs []day14Reindeer, sec uint) uint {
	if len(rs) == 0 {
		return 0
	}
	scores := make([]uint, len(rs))
	for t := uint(1); t <= sec; t++ {
		lead := uint(0)
		for _, r := range rs {
			lead = max(lead, day14Distance(r, t))
		}
		for i, r := range rs {
			if day14Distance(r, t) == lead {
				scores[i]++
			}
		}
	}
	best := uint(0)
	for _, s := range scores {
		best = max(best, s)
	}
	return best
}

