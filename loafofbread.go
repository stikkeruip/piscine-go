package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	finStr := ""
	count := 0
	for _, r := range str {
		if r == ' ' {
			continue
		}
		if count == 5 {
			finStr += " "
			count = 0
		}
		finStr += string(r)
		count++
	}
	if len(finStr) > 0 && finStr[len(finStr)-1] == ' ' {
		finStr = finStr[:len(finStr)-1]
	}
	return finStr + "\n"
}
