package adventofcode2015

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type spellID int

const (
	magicMissile spellID = iota
	drain
	shield
	poison
	recharge
)

type day22Player struct {
	hitPoints int
	damage    int
}

type Day22Puzzle struct {
	boss day22Player
}

func NewDay22(lines []string) (Day22Puzzle, error) {
	if len(lines) != 3 {
		return Day22Puzzle{}, fmt.Errorf("invalid input")
	}
	hitPoints, err := day22ParseStat(lines[0], "Hit Points:")
	if err != nil {
		return Day22Puzzle{}, err
	}
	damage, err := day22ParseStat(lines[1], "Damage:")
	if err != nil {
		return Day22Puzzle{}, err
	}
	return Day22Puzzle{
		boss: day22Player{hitPoints: hitPoints, damage: damage},
	}, nil
}

func day22ParseStat(line, prefix string) (int, error) {
	if !strings.HasPrefix(line, prefix) {
		return 0, fmt.Errorf("invalid line %q", line)
	}
	n, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, prefix)))
	if err != nil {
		return 0, err
	}
	return n, nil
}

const (
	shieldTimer   = 6
	poisonTimer   = 6
	rechargeTimer = 5
)

// Day22 solves day 22 for the selected part.
func Day22(puzzle Day22Puzzle, part1 bool) uint {
	return uint(day22(!part1, puzzle.boss))
}

func day22(hardMode bool, boss day22Player) int {
	const startMana = 500
	const startHP = 50

	start := day22State{
		playerHP:    startHP,
		playerMana:  startMana,
		bossHP:      boss.hitPoints,
		bossDamage:  boss.damage,
		playerTurn:  true,
		shieldTimer: 0,
		poisonTimer: 0,
		rechTimer:   0,
		spent:       0,
	}

	pq := make(day22PQ, 0, 256)
	heap.Push(&pq, start)

	bestWin := int(^uint(0) >> 1)
	seen := make(map[day22Key]int, 4096)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(day22State)
		if cur.spent >= bestWin {
			continue
		}

		key := cur.key()
		if best, ok := seen[key]; ok && best <= cur.spent {
			continue
		}
		seen[key] = cur.spent

		if hardMode && cur.playerTurn {
			cur.playerHP--
			if cur.playerHP <= 0 {
				continue
			}
		}

		cur.tickEffects()
		if cur.bossHP <= 0 {
			if cur.spent < bestWin {
				bestWin = cur.spent
			}
			continue
		}

		if cur.playerTurn {
			for spell := magicMissile; spell <= recharge; spell++ {
				next, ok := cur.cast(spell)
				if !ok || next.spent >= bestWin {
					continue
				}
				if next.bossHP <= 0 {
					if next.spent < bestWin {
						bestWin = next.spent
					}
					continue
				}
				next.playerTurn = false
				heap.Push(&pq, next)
			}
			continue
		}

		armor := 0
		if cur.shieldTimer > 0 {
			armor = 7
		}
		damage := cur.bossDamage - armor
		if damage < 1 {
			damage = 1
		}
		cur.playerHP -= damage
		if cur.playerHP <= 0 {
			continue
		}
		cur.playerTurn = true
		heap.Push(&pq, cur)
	}

	return bestWin
}

type day22State struct {
	playerHP   int
	playerMana int
	bossHP     int
	bossDamage int
	spent      int

	shieldTimer int
	poisonTimer int
	rechTimer   int
	playerTurn  bool
}

type day22Key struct {
	playerHP    int16
	playerMana  int16
	bossHP      int16
	shieldTimer int8
	poisonTimer int8
	rechTimer   int8
	playerTurn  uint8
}

func (a day22State) key() day22Key {
	playerTurn := uint8(0)
	if a.playerTurn {
		playerTurn = 1
	}
	return day22Key{
		playerHP:    int16(a.playerHP),
		playerMana:  int16(a.playerMana),
		bossHP:      int16(a.bossHP),
		shieldTimer: int8(a.shieldTimer),
		poisonTimer: int8(a.poisonTimer),
		rechTimer:   int8(a.rechTimer),
		playerTurn:  playerTurn,
	}
}

func (a *day22State) tickEffects() {
	if a.shieldTimer > 0 {
		a.shieldTimer--
	}
	if a.poisonTimer > 0 {
		a.bossHP -= 3
		a.poisonTimer--
	}
	if a.rechTimer > 0 {
		a.playerMana += 101
		a.rechTimer--
	}
}

func (a day22State) cast(spell spellID) (day22State, bool) {
	next := a
	switch spell {
	case magicMissile:
		if next.playerMana < 53 {
			return day22State{}, false
		}
		next.playerMana -= 53
		next.spent += 53
		next.bossHP -= 4
		return next, true
	case drain:
		if next.playerMana < 73 {
			return day22State{}, false
		}
		next.playerMana -= 73
		next.spent += 73
		next.bossHP -= 2
		next.playerHP += 2
		return next, true
	case shield:
		if next.playerMana < 113 || next.shieldTimer > 0 {
			return day22State{}, false
		}
		next.playerMana -= 113
		next.spent += 113
		next.shieldTimer = shieldTimer
		return next, true
	case poison:
		if next.playerMana < 173 || next.poisonTimer > 0 {
			return day22State{}, false
		}
		next.playerMana -= 173
		next.spent += 173
		next.poisonTimer = poisonTimer
		return next, true
	case recharge:
		if next.playerMana < 229 || next.rechTimer > 0 {
			return day22State{}, false
		}
		next.playerMana -= 229
		next.spent += 229
		next.rechTimer = rechargeTimer
		return next, true
	default:
		return day22State{}, false
	}
}

type day22PQ []day22State

func (a day22PQ) Len() int { return len(a) }

func (a day22PQ) Less(i, j int) bool {
	return a[i].spent < a[j].spent
}

func (a day22PQ) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a *day22PQ) Push(x any) {
	*a = append(*a, x.(day22State))
}

func (a *day22PQ) Pop() any {
	old := *a
	n := len(old)
	last := old[n-1]
	*a = old[:n-1]
	return last
}
