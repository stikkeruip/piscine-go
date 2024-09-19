package piscine

func Rot14(s string) string {
	runes := []rune(s) // Convert the input string to a slice of runes.

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + (r-'a'+14)%26 // Rotate within lowercase letters.
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + (r-'A'+14)%26 // Rotate within uppercase letters.
		}
	}

	return string(runes) // Convert the modified runes back to a string.
}
