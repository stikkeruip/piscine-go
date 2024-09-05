package piscine

func BasicAtoi(s string) int {
	rns := []rune(s)
	result := 0
	for i := 0; i < len(rns); i++ {
		result = result*10 + int(rns[i]-'0')
	}
	return result
}
