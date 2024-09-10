package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)

	for _, i := range runes {
		if (i > 'z' || i < 'a') && (i > 'Z' || i < 'A') && (i > '9' || i < '0') {
			return false
		}
	}
	return true
}
