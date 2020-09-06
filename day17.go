package adventofcode2015

// Day17 returns number of combinations that storage can be fit into capacities.
func Day17(storage uint, capacities []uint) uint {
	var n uint
	sum := func(ns []uint) uint {
		var m uint
		for _, n := range ns {
			m += n
		}
		return m
	}

	for _, perm := range PowerSet(capacities) {
		if sum(perm) == storage {
			n++
		}
	}
	return n
}

func PowerSet(original []uint) [][]uint {
	powerSetSize := 1 << len(original)
	result := make([][]uint, 0, powerSetSize)

	var index int
	for index < powerSetSize {
		var subSet []uint

		for j, elem := range original {
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		result = append(result, subSet)
		index++
	}
	return result
}
