package piscine

func IsUpper(s string) bool {
	runes := []rune(s)

	for _, i := range runes {
		if i > 'Z' || i < 'A' {
			return false
		}
	}
	return true
}
