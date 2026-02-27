package adventofcode2015

import (
	"fmt"
	"io"
	"math"
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
	const (
		startMana = 500
		maxRates  = 13
	)
	minSpent := math.MaxInt16

	for digits := NewBase5(maxRates); len(digits.Buf) <= maxRates; digits.Inc() {
		players := [...]day22Player{
			{hitPoints: 50, armor: 0, mana: startMana},
			boss,
		}
		i := -1
		f := func() spellID {
			i++
			return spellID(digits.Buf[i])
		}
		g := newWizardSimulator(players, f, io.Discard)
		// g := newWizardSimulator(players, f, os.Stdout)
		playerFailed := false
		spentTooMuch := false
		for !(playerFailed || spentTooMuch || g.gameOver()) {
			err := g.step(hardMode)
			if err != nil {
				playerFailed = true
			}
			// no need to keep going if we already spent more than
			// current best
			if g.players[0].spent > minSpent {
				spentTooMuch = true
			}
		}
		if playerFailed || spentTooMuch || g.playerLost() {
			continue
		}
		if g.players[0].spent < minSpent {
			minSpent = g.players[0].spent
		}
	}
	return minSpent
}
