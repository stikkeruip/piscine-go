package piscine

func Sqrt(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	i, result := 1, 1

	for result <= n {
		i++
		result = i * i
	}
	if (i-1)*(i-1) != n {
		return 0
	}
	return i - 1
}
