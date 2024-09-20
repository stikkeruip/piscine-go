package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	// Remove spaces and get the letters
	letters := ""
	for _, r := range str {
		if r != ' ' {
			letters += string(r)
		}
	}
	if len(letters) < 5 {
		return "Invalid Input\n"
	}
	// Initialize variables
	finStr := ""
	count := 0
	for i, r := range letters {
		finStr += string(r)
		count++
		// Add a space after every 5 letters, except after the last group
		if count == 5 && i != len(letters)-1 {
			finStr += " "
			count = 0
		}
	}
	return finStr + "\n"
}
