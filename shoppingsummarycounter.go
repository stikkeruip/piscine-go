package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	receipt := make(map[string]int)
	word := ""
	for _, r := range str {
		if r == ' ' {
			receipt[word]++
			word = ""
			continue
		}
		word += string(r)
	}
	receipt[word]++
	return receipt
}
