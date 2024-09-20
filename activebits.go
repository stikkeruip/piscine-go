package piscine

func ActiveBits(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}
	return 1 + ActiveBits(n/2)
}
