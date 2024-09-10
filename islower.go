package piscine

func IsLower(s string) bool {
	runes := []rune(s)

	for _, i := range runes {
		if i > 'z' || i < 'a' {
			return false
		}
	}
	return true
}
