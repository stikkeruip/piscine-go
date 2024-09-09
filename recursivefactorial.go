package piscine

func RecursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return RecursiveFactorial(n-1) * n
}
