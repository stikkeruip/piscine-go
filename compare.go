package piscine

func Compare(a string, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
