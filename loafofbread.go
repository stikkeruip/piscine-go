package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}

	// Remove spaces to get all letters
	letters := ""
	for _, r := range str {
		if r != ' ' {
			letters += string(r)
		}
	}

	// Check if there are at least 5 letters
	if len(letters) < 5 {
		return "Invalid Input\n"
	}

	// Build the final string with spaces after every 5 letters
	finStr := ""
	for i, r := range letters {
		finStr += string(r)
		if (i+1)%5 == 0 && i != len(letters)-1 {
			finStr += " "
		}
	}

	return finStr + "\n"
}
