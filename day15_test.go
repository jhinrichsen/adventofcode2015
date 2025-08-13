package adventofcode2015

import (
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestDay15NewIngredient(t *testing.T) {
	const s = "Butterscotch: capacity -1, durability -2, flavor 6, " +
		"texture 3, calories 8"
	i, err := NewIngredient(s)
	if err != nil {
		t.Fatal(err)
	}
	if i.Name != "Butterscotch" {
		t.Fatalf("want Butterscotch but got %q", i.Name)
	}
	if i.Properties["capacity"] != -1 {
		t.Fatalf("capacity is broke")
	}
	if i.Properties["calories"] != 8 {
		t.Fatal("calories is broke")
	}
}

func ingredients(filename string) ([]Ingredient, error) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		return nil, err
	}
	var is []Ingredient
	for _, line := range lines {
		i, err := NewIngredient(line)
		if err != nil {
			return nil, err
		}
		is = append(is, i)
	}
	return is, nil
}

func TestDay15Example(t *testing.T) {
	const want = 62_842_880
	is, err := ingredients(exampleFilename(15))
	if err != nil {
		t.Fatal(err)
	}
	if len(is) != 2 {
		t.Fatalf("want butterscotch and cinnamon but got %q", is)
	}
	// Butterscotch
	var c Cookie
	c = append(c, Serving{is[0], 44}) // butterscotch
	c = append(c, Serving{is[1], 56}) // cinnamon
	got := c.score()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15ProveExample(t *testing.T) {
	const want = 62_842_880
	is, err := ingredients(exampleFilename(15))
	if err != nil {
		t.Fatal(err)
	}
	cookie := Day15Part1(is)
	got := cookie.score()
	die(want, got, cookie, t)
}

func TestDay15Part1(t *testing.T) {
	const want = 13_882_464
	is, err := ingredients(filename(15))
	if err != nil {
		t.Fatal(err)
	}
	cookie := Day15Part1(is)
	got := cookie.score()
	die(want, got, cookie, t)
}

func TestDay15Part2(t *testing.T) {
	const want = 11_171_160
	is, err := ingredients(filename(15))
	if err != nil {
		t.Fatal(err)
	}
	cookie := Day15Part2(is)
	got := cookie.score()
	die(want, got, cookie, t)
}

func die(want, got uint, cookie Cookie, t *testing.T) {
	if want != got {
		// long numbers, print using thousand separator which fmt does
		// not support
		p := message.NewPrinter(language.English)
		s := p.Sprintf("want %d but got %d: champ: %+v", want, got, cookie)
		t.Fatalf("%s", s)
	}
}
