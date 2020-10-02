package adventofcode2015

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
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
		fmt.Fprintf(a.progress,
			"Player casts %s, dealing 4 damage.\n",
			a.spells[s].name)
	case drain:
		a.players[0].hitPoints += 2
		a.players[1].hitPoints -= 2
		fmt.Fprintf(a.progress,
			"Player casts Drain, dealing 2 damage, "+
				"and healing 2 hit points.\n")

	case shield:
		a.decreaseTimer(s)
		fmt.Fprintf(a.progress, "Shield's timer is now %d.\n",
			a.spells[s].timer)
		if a.spells[s].timer == 0 {
			fmt.Fprintf(a.progress, "Shield wears off, "+
				"decreasing armor by 7.\n")
		}
		a.resetTimer(s)
	case poison:
		a.players[1].hitPoints -= 3
		if a.bossLost() {
			fmt.Fprintf(a.progress,
				"Poison deals 3 damage. "+
					"This kills the boss, "+
					"and the player wins.\n")
		} else {
			a.decreaseTimer(s)
			fmt.Fprintf(a.progress,
				"Poison deals 3 damage; its timer is now %d.\n",
				a.spells[s].timer)
			a.resetTimer(s)
		}
	case recharge:
		a.players[0].mana += 101
		a.decreaseTimer(s)
		fmt.Fprintf(a.progress, "Recharge provides 101 mana; "+
			"its timer is now %d.\n",
			a.spells[s].timer)
		if a.spells[s].timer == 0 {
			fmt.Fprintf(a.progress, "Recharge wears off.\n")
		}
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

// pointOrPoints returns "point" or "points", depending on n.
func pointOrPoints(n int) string {
	if n == 1 {
		return "point"
	}
	return "points"
}

func (a wizardSimulator) logTurn() {
	var s string
	if a.turnIdx == 0 {
		s = "Player"
	} else {
		s = "Boss"
	}
	fmt.Fprintf(a.progress, "-- %s turn --\n", s)
	fmt.Fprintf(a.progress, "- Player has %d hit %s, "+
		"%d armor, %d mana\n",
		a.players[0].hitPoints,
		pointOrPoints(a.players[0].hitPoints),
		a.totalArmor(),
		a.players[0].mana)
	fmt.Fprintf(a.progress, "- Boss has %d hit %s\n",
		a.players[1].hitPoints,
		pointOrPoints(a.players[1].hitPoints))
}

// step returns error if cannot afford spell, or if spell is active.
func (a *wizardSimulator) step(hardMode bool) error {
	a.logTurn()
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
			if spell == shield {
				fmt.Fprintf(a.progress, "Player casts Shield, "+
					"increasing armor by 7.\n")
			} else {
				fmt.Fprintf(a.progress, "Player casts %s.\n",
					a.spells[spell].name)
			}
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

// Day22Part1 returns the least amount of mana you can spend and still win the
// fight.
// Of all the AOC solutions, this one is the most ugly one. And the slowest.
func Day22Part1() int {
	return day22(false)
}

// Day22Part2 runs day 22 in hard mode.
func Day22Part2() int {
	return day22(true)
}

func day22(hardMode bool) int {
	const (
		startMana = 500
		maxRates  = 13
	)
	minSpent := math.MaxInt16
	digits := NewBase5(maxRates)

	for digits = NewBase5(maxRates); len(digits.Buf) <= maxRates; digits.Inc() {
		players := [...]day22Player{
			{hitPoints: 50, armor: 0, mana: startMana},
			{hitPoints: 58, damage: 9},
		}
		i := -1
		f := func() spellID {
			i++
			return spellID(digits.Buf[i])
		}
		g := newWizardSimulator(players, f, ioutil.Discard)
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
