package piscine

func Capitalize(s string) string {
	runes := []rune(s)

	inWord := false

	for i, r := range runes {
		if inWord == true && (r < 'a' || r > 'z') {
			inWord = false
		}

		if inWord == false && (r >= 'a' && r <= 'z') {
			runes[i] = r - 32
			inWord = true
		}
		if inWord == false && (r >= 'A' && r <= 'Z') {
			inWord = true
		}
	}

	return string(runes)
}
