package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}

	// Count the number of letters (excluding spaces)
	letterCount := 0
	for _, r := range str {
		if r != ' ' {
			letterCount++
		}
	}

	// If less than 5 letters, return "Invalid Output\n"
	if letterCount < 5 {
		return "Invalid Output\n"
	}

	// Process the input string
	finStr := ""
	lettersCollected := 0
	strRunes := []rune(str) // Convert string to rune slice for proper indexing
	i := 0
	for i < len(strRunes) {
		r := strRunes[i]
		if r != ' ' {
			finStr += string(r)
			lettersCollected++
			if lettersCollected == 5 {
				// Skip the next character in the input string
				i++
				lettersCollected = 0
				// Add a space if not at the end
				if i < len(strRunes) {
					finStr += " "
				}
				continue
			}
		}
		i++
	}

	return finStr + "\n"
}
