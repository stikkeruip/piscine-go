package piscine

func Map(f func(int) bool, a []int) []bool {
	var slice []bool

	for _, n := range a {
		slice = append(slice, f(n))
	}

	return slice
}
