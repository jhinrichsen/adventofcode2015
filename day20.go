package adventofcode2015

// Input for part 1.
const Input = 33_100_000

// presents returns the number of presents all elfs visiting a house have left.
func presents(houseno uint) uint {
	if houseno == 0 {
		return 0
	}
	var presents uint
	for elf := uint(1); elf <= houseno; elf++ {
		if houseno%elf == 0 {
			// yes, elf visits this house
			presents += elf * 10
		}
	}
	return presents
}

// Day20 returns lowest house no that gets at least n presents.
// Brute forcing without a hint takes 10 min on 2019 Macbook Pro 16".
func Day20(n uint) uint {
	hint := uint(500_000)
	houseno := hint
	for {
		houseno++
		p := presents(houseno)
		// log.Printf("%d: %d\n", houseno, p)
		if p >= n {
			break
		}
	}
	return houseno
}
