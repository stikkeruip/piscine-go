package piscine

func Unmatch(a []int) int {
	seen := make(map[int]int)

	for _, n := range a {
		seen[n]++
	}
	for key, value := range seen {
		if value%2 != 0 {
			return key
		}
	}
	return -1
}
