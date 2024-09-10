package piscine

func ToLower(s string) string {
	runes := []rune(s)

	for index, i := range runes {
		if i <= 'Z' && i >= 'A' {
			runes[index] = i + 32
		}
	}
	return string(runes)
}
