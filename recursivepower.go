package piscine

func RecursivePower(n int, p int) int {
	if p < 0 {
		return 0
	}

	if p == 0 {
		return 1
	}

	return n * RecursivePower(n, p-1)
}
