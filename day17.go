package adventofcode2015

// Day17Part1 returns number of combinations that storage can be fit into
// capacities.
func Day17Part1(storage uint, capacities []uint) uint {
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

// PowerSet returns all combinations (including empty one) for given array.
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

// Day17Part2 returns number of combinations that storage can be fit into
// capacities.
func Day17Part2(storage uint, capacities []uint) uint {
	// maps holds number of entriees of length n
	m := make(map[uint]uint)
	sum := func(ns []uint) uint {
		var m uint
		for _, n := range ns {
			m += n
		}
		return m
	}

	for _, perm := range PowerSet(capacities) {
		if sum(perm) == storage {
			l := uint(len(perm))
			m[l]++
			// fmt.Printf("m[%d]=%d\n", l, m[l])
		}
	}

	// find smallest existing index, return number of occurrences
	for i := uint(1); i < storage; i++ {
		if n, ok := m[i]; ok {
			return n
		}
	}
	// this is unreachable code, but instead of panic() we
	return 0
}
