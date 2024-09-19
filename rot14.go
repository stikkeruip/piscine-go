package piscine

func Rot14(s string) string {
	runes := []rune(s)

	for _, r := range s {
		runes = append(runes, r+14)
	}

	return string(runes)
}
