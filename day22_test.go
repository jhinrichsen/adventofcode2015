package adventofcode2015

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// numbered turns multiple strings into one, separated by newline, and prefixed with the line number.
func numbered(ss []string) string {
	var sb strings.Builder
	for i, s := range ss {
		sb.WriteString(fmt.Sprintf("%3d %s\n", i, s))
	}
	return sb.String()
}

func testWizardSimulator(players [2]day22Player,
	spells []spellID, example uint) error {
	filename := fmt.Sprintf("testdata/day22_example%d.txt", example)
	wantLines, err := linesFromFilename(filename)
	if err != nil {
		return err
	}

	var w bytes.Buffer
	i := -1
	g := newWizardSimulator(players, func() spellID {
		i++
		return spells[i]
	}, &w)
	for !g.gameOver() {
		err := g.step(false)
		if err != nil {
			io.Copy(os.Stdout, &w)
			return err
		}
	}
	gotLines, err := linesFromReader(&w)
	if err != nil {
		return err
	}
	for i := range wantLines {
		if wantLines[i] != gotLines[i] {
			s := numbered(gotLines)
			return fmt.Errorf("line %d: want %q but got %q\n%+v", i,
				wantLines[i], gotLines[i], s)
		}
	}
	return nil
}

func TestDay22Example1(t *testing.T) {
	players := [...]day22Player{
		{hitPoints: 10, armor: 0, mana: 250},
		{hitPoints: 13, damage: 8},
	}
	spells := []spellID{
		poison,
		magicMissile,
	}
	err := testWizardSimulator(players, spells, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDay22Example2(t *testing.T) {
	players := [...]day22Player{
		{hitPoints: 10, armor: 0, mana: 250},
		{hitPoints: 14, damage: 8},
	}
	spells := []spellID{
		recharge,
		shield,
		drain,
		poison,
		magicMissile,
	}
	err := testWizardSimulator(players, spells, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDay22CastActiveSpell(t *testing.T) {
	players := [...]day22Player{
		{hitPoints: 10, armor: 0, mana: 500},
		{hitPoints: 14, damage: 8},
	}
	spells := []spellID{
		recharge,
		recharge,
	}
	err := testWizardSimulator(players, spells, 2)
	if err == nil {
		t.Fatal("want error but got none")
	}
}

func TestDay22Part1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long running testcase for day 22")
	}
	const want = 1269
	got := Day22Part1()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long running testcase for day 22")
	}
	const want = 1309
	got := Day22Part2()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
