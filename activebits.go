package piscine

func ActiveBits(n int) int {
	if n == 0 {
		return 0
	}
	return n%2 + ActiveBits(n/2)
}
