package piscine

func IterativePower(n int, p int) int {
	if n == 0 || p < 0 {
		return 0
	}

	if p == 0 {
		return 1
	}

	result := n

	for i := 1; i < p; i++ {
		result *= n
	}
	return result
}
