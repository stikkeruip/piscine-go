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
	if finStr[len(finStr)-1] == ' ' {
		finStr = finStr[:len(finStr)-1]
	}
	return finStr + "\n"
}
