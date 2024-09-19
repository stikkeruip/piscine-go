package piscine

func Unmatch(a []int) int {
	seen := make(map[int]int)

	for _, n := range a {
		seen[n]++
	}

	unmatched := -1
	for key, value := range seen {
		if value%2 != 0 {
			if key > unmatched {
				unmatched = key
			}
		}
	}
	return unmatched
}
