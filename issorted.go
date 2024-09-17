package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 1 {
		return true
	}

	isAscending := true
	isDescending := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			isDescending = false
		} else {
			isAscending = false
		}
	}
	return !(isAscending == isDescending)
}
