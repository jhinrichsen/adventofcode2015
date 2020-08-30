package adventofcode2015

import "testing"

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

func TestDay15Example(t *testing.T) {
	const want = 62842880
	lines, err := linesFromFilename(exampleFilename(15))
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("want butterscotch and cinnamon but got %q", lines)
	}
	var is []Ingredient
	for _, line := range lines {
		i, err := NewIngredient(line)
		if err != nil {
			t.Fatal(err)
		}
		is = append(is, i)
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
