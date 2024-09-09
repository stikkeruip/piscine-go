package piscine

func IterativeFactorial(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	fac := 1

	for i := 1; i <= n; i++ {
		fac = fac * i
	}

	return fac
}
