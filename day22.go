package adventofcode2015

import (
	"container/heap"
	"fmt"
	"io"
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

type spell struct {
	name   string
	mana   int
	timer  int
	active bool
}

// immediate indicates non-effects (no timer)
const immediate = -1

func (a spell) isEffect() bool {
	return a.timer != immediate
}

type day22Player struct {
	hitPoints int
	armor     int
	mana      int
	spent     int // recharge may not affect mana, so need a separate counter
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

// wizardSimulator holds data for one game.
type wizardSimulator struct {
	players  [2]day22Player
	caster   func() spellID
	spells   map[spellID]spell
	turnIdx  int
	progress io.Writer
}

// NewWizardSimulator returns a new wizard game.
func newWizardSimulator(players [2]day22Player,
	f func() spellID,
	progress io.Writer) wizardSimulator {

	return wizardSimulator{players, f, allSpells(), 0, progress}
}

const (
	shieldTimer   = 6
	poisonTimer   = 6
	rechargeTimer = 5
)

func allSpells() map[spellID]spell {
	spells := make(map[spellID]spell)
	spells[magicMissile] = spell{name: "Magic Missile", mana: 53, timer: immediate}
	spells[drain] = spell{name: "Drain", mana: 73, timer: immediate}
	spells[shield] = spell{name: "Shield", mana: 113, timer: shieldTimer}
	spells[poison] = spell{name: "Poison", mana: 173, timer: poisonTimer}
	spells[recharge] = spell{name: "Recharge", mana: 229, timer: rechargeTimer}
	return spells
}

// applyEffects executes spells (both immediate and effects)
func (a *wizardSimulator) applyEffects() {
	// in Go, map iteration order is undefined. In order to get consistent
	// progress logging, do ordered range.
	keys := []spellID{shield, poison, recharge}
	for _, k := range keys {
		v := a.spells[k]
		if v.active {
			a.apply(k)
		}
	}
}

func (a *wizardSimulator) apply(s spellID) {
	switch s {
	case magicMissile:
		a.players[1].hitPoints -= 4
	case drain:
		a.players[0].hitPoints += 2
		a.players[1].hitPoints -= 2

	case shield:
		a.decreaseTimer(s)
		a.resetTimer(s)
	case poison:
		a.players[1].hitPoints -= 3
		if a.bossLost() {
		} else {
			a.decreaseTimer(s)
			a.resetTimer(s)
		}
	case recharge:
		a.players[0].mana += 101
		a.decreaseTimer(s)
		a.resetTimer(s)
	}
}

func (a *wizardSimulator) decreaseTimer(spell spellID) {
	s := a.spells[spell]
	s.timer--
	a.spells[spell] = s
}

func (a *wizardSimulator) resetTimer(spell spellID) {
	s := a.spells[spell]
	if s.timer == 0 {
		s.active = false
		s.timer = resetTimer(spell)
	}
	a.spells[spell] = s
}

func resetTimer(spell spellID) int {
	switch spell {
	case poison:
		return poisonTimer
	case shield:
		return shieldTimer
	case recharge:
		return rechargeTimer
	}
	return immediate
}

func (a wizardSimulator) gameOver() bool {
	return a.playerLost() || a.bossLost()
}

func (a wizardSimulator) playerLost() bool {
	return a.players[0].hitPoints < 1
}

func (a wizardSimulator) canAfford(ID spellID) bool {
	return a.players[0].mana >= a.spells[ID].mana
}

func (a wizardSimulator) bossLost() bool {
	return a.players[1].hitPoints < 1
}

func (a wizardSimulator) totalArmor() int {
	if a.spells[shield].active {
		return a.players[0].armor + 7
	}
	return a.players[0].armor
}

// // step returns error if cannot afford spell, or if spell is active.
func (a *wizardSimulator) step(hardMode bool) error {
	if hardMode {
		a.players[0].hitPoints--
		if a.playerLost() {
			return nil
		}
	}
	a.applyEffects()
	if a.bossLost() {
		return nil
	}

	if a.turnIdx == 0 {
		spell := a.caster()
		if !a.canAfford(spell) {
			return fmt.Errorf("cannot afford %q: have %d but need %d",
				spell, a.players[0].mana, a.spells[spell].mana)
		}
		if a.spells[spell].active {
			return fmt.Errorf("cannot cast active spell %+v: %+v", spell, a.spells[spell])
		}
		if a.spells[spell].isEffect() {
			sp := a.spells[spell]
			sp.active = true
			a.spells[spell] = sp
		} else {
			a.apply(spell)
		}

		// adjust player mana
		a.players[0].mana -= a.spells[spell].mana
		a.players[0].spent += a.spells[spell].mana
	} else {
		damage := a.players[a.turnIdx].damage -
			a.totalArmor()
		if damage < 1 {
			damage = 1
		}
		a.players[1-a.turnIdx].hitPoints -= damage
		var msg string
		if a.spells[shield].active {
			msg = fmt.Sprintf("%d - 7 = %d", damage+7, damage)
		} else {
			msg = fmt.Sprintf("%d", damage)
		}
		fmt.Fprintf(a.progress, "Boss attacks for %s damage.\n", msg)
	}
	fmt.Fprintf(a.progress, "\n")
	a.turn()
	return nil
}

func (a *wizardSimulator) turn() {
	a.turnIdx = 1 - a.turnIdx
}

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
