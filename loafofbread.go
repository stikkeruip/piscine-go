package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 || str == "" {
		return "Invalid Input\n"
	}
	finStr := ""
	count := 0
	for _, r := range str {
		if count == 5 {
			finStr += " "
			count = 0
			continue
		}
		if r == ' ' {
			continue
		}
		finStr += string(r)
		count++
	}
	return finStr + "\n"
}
