package piscine

func LoafOfBread(str string) string {
	// Remove spaces to get all letters
	letters := ""
	for _, r := range str {
		if r != ' ' {
			letters += string(r)
		}
	}

	// If no letters, return "\n"
	if len(letters) == 0 {
		return "\n"
	}

	// If less than 5 letters, return "Invalid Output\n"
	if len(letters) < 5 {
		return "Invalid Output\n"
	}

	// Process the input string
	finStr := ""
	lettersCollected := 0
	skipNext := false
	strRunes := []rune(str) // Convert string to rune slice for proper indexing
	i := 0
	for i < len(strRunes) {
		if skipNext {
			skipNext = false
			i++
			continue
		}

		r := strRunes[i]
		if r != ' ' {
			finStr += string(r)
			lettersCollected++
			if lettersCollected == 5 {
				lettersCollected = 0
				skipNext = true // Skip next character in the input string
				// Add a space if there are more letters ahead (after skipping next character)
				if hasMoreLettersAfter(i, strRunes) {
					finStr += " "
				}
			}
		}
		i++
	}

	return finStr + "\n"
}

// Helper function to check if there are more letters after the current position
func hasMoreLettersAfter(pos int, strRunes []rune) bool {
	// Skip the next character in the input string
	pos += 2
	for j := pos; j < len(strRunes); j++ {
		if strRunes[j] != ' ' {
			return true
		}
	}
	return false
}
