package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)

	result := 0

	sign := 1

	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			result = result*10 + int(runes[i]-'0')
		}
		if runes[i] == '-' && result == 0 {
			sign = -1
		}
	}

	return result * sign
}
