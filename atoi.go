package piscine

func Atoi(s string) int {
	rns := []rune(s)
	sign := 1
	start := 0
	result := 0

	if rns[0] == '+' {
		start = 1
	}

	if rns[0] == '-' {
		start = 1
		sign = -1
	}

	for i := start; i < len(rns); i++ {
		if rns[i] >= '0' && rns[i] <= '9' {
			result = result*10 + int(rns[i]-'0')
		} else {
			result = 0
			break
		}
	}
	return result * sign
}
