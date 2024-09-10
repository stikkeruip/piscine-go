package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)

	for _, i := range runes {
		if i > '9' || i < '0' {
			return false
		}
	}
	return true
}
