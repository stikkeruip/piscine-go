package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	newStr := ""

	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			newStr += string(str[i])
		}
	}
	return newStr + "\n"
}
