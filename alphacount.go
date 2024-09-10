package piscine

func AlphaCount(s string) int {
	runes := []rune(s)

	count := 0

	for _, i := range runes {
		if i >= 'a' && i <= 'Z' {
			count++
		}
	}
	return count
}
