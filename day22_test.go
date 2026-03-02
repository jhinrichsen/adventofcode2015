package adventofcode2015

import (
	"errors"
	"testing"
)

func TestDay22Example1(t *testing.T) {
	start := day22State{
		playerHP:   10,
		playerMana: 250,
		bossHP:     13,
		bossDamage: 8,
	}
	got, err := day22Replay(start, []spellID{poison, magicMissile})
	if err != nil {
		t.Fatal(err)
	}
	if got.bossHP > 0 {
		t.Fatalf("want boss dead but hp=%d", got.bossHP)
	}
	if got.spent != 226 {
		t.Fatalf("want %d but got %d", 226, got.spent)
	}
}

func TestDay22Example2(t *testing.T) {
	start := day22State{
		playerHP:   10,
		playerMana: 250,
		bossHP:     14,
		bossDamage: 8,
	}
	got, err := day22Replay(start, []spellID{recharge, shield, drain, poison, magicMissile})
	if err != nil {
		t.Fatal(err)
	}
	if got.bossHP > 0 {
		t.Fatalf("want boss dead but hp=%d", got.bossHP)
	}
	if got.spent != 641 {
		t.Fatalf("want %d but got %d", 641, got.spent)
	}
}

func TestDay22Part1(t *testing.T) {
	testWithParser(t, 22, filename, true, NewDay22, Day22, uint(1_269))
}

func TestDay22Part2(t *testing.T) {
	testWithParser(t, 22, filename, false, NewDay22, Day22, uint(1_309))
}

func BenchmarkDay22Part1(b *testing.B) {
	benchWithParser(b, 22, true, NewDay22, Day22)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchWithParser(b, 22, false, NewDay22, Day22)
}

func day22Replay(start day22State, spells []spellID) (day22State, error) {
	s := start
	i := 0
	for s.playerHP > 0 && s.bossHP > 0 {
		s.tickEffects()
		if s.bossHP <= 0 {
			return s, nil
		}
		if i >= len(spells) {
			return s, errors.New("not enough spells in sequence")
		}
		next, ok := s.cast(spells[i])
		if !ok {
			return s, errors.New("invalid spell cast")
		}
		s = next
		i++
		if s.bossHP <= 0 {
			return s, nil
		}

		s.tickEffects()
		if s.bossHP <= 0 {
			return s, nil
		}
		armor := 0
		if s.shieldTimer > 0 {
			armor = 7
		}
		damage := s.bossDamage - armor
		if damage < 1 {
			damage = 1
		}
		s.playerHP -= damage
	}
	return s, nil
}
