package piscine

func IterativeFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return IterativeFactorial(n-1) * n
}
