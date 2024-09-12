package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := baseToDecimal(nbr, baseFrom)

	return decimalToBase(decimalValue, baseTo)
}

func baseToDecimal(nbr, baseFrom string) int {
	decimalValue := 0
	baseLength := len(baseFrom)

	baseMap := make(map[rune]int)
	for i, ch := range baseFrom {
		baseMap[ch] = i
	}

	for _, ch := range nbr {
		decimalValue = decimalValue*baseLength + baseMap[ch]
	}
	return decimalValue
}

func decimalToBase(decimalValue int, baseTo string) string {
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	baseLength := len(baseTo)
	result := ""

	for decimalValue > 0 {
		remainder := decimalValue % baseLength
		result = string(baseTo[remainder]) + result
		decimalValue /= baseLength
	}
	return result
}
