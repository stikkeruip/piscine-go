package piscine

func ToUpper(s string) string {
	runes := []rune(s)

	for index, i := range runes {
		if i <= 'z' && i >= 'a' {
			runes[index] = i - 32
		}
	}
	return string(runes)
}
