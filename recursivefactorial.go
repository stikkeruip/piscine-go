package piscine

func RecursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	if n >= 21 || n < 0 {
		return 0
	}

	return RecursiveFactorial(n-1) * n
}
