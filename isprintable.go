package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)

	for _, i := range runes {
		if i > '~' || i < ' ' {
			return false
		}
	}
	return true
}
