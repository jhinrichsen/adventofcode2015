package adventofcode2015

// InputDay20 for part 1.
const InputDay20 = 33_100_000

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
// A memoized version of Euler's recursive Sigma function brings the calculation
// down to 48 seconds.
func Day20(n uint) uint {
	for houseno := uint(1); ; houseno++ {
		s := SigmaMemoized(houseno)
		if s*10 > InputDay20 {
			return houseno
		}
	}
}

// day20Champ is the highest ranking algo.
func day20Champ(target int) int {
	houses := make([]int, target/10+1)
	for elf := 1; elf < len(houses); elf++ {
		for house := elf; house < len(houses); house += elf {
			houses[house] += elf * 11
		}
	}
	for house := 1; house < len(houses); house++ {
		if houses[house] > target {
			return house
		}
	}
	return -1
}

// day20MyChamp is the highest ranking algo.
func day20MyChamp(target int) int {
	houses := make([]int, target/10+1)
	for elf := 1; elf < len(houses); elf++ {
		for house := elf; house < len(houses); house += elf {
			houses[house] += elf * 11
			if houses[house] > target {
				return house
			}
		}
	}
	return -1
}
