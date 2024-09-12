package piscine

func Split(s, sep string) []string {
	var result []string
	word := ""
	for i := 0; i < len(s); i++ {
		if len(s[i:]) >= len(sep) && s[i:i+len(sep)] == sep {
			if word != "" {
				result = append(result, word)
				word = ""
			}
			i += len(sep) - 1
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
