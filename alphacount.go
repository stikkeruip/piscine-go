package piscine

func AlphaCount(s string) int {
	runes := []string(s)

	count := 0

	for _, i ; range runes {
		if (i >= 'a' && i <= 'Z') {
			count++
		}
	}
	return count
}