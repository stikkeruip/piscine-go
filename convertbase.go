package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := binaryToDecimal(nbr, baseFrom)

	return decimalToBase(decimalValue, baseTo)
}

func binaryToDecimal(nbr, baseFrom string) int {
	decimalValue := 0
	baseLength := len(baseFrom)

	for _, ch := range nbr {
		decimalValue = decimalValue*baseLength + int(ch-'0')
	}
	return decimalValue
}

func decimalToBase(decimalValue int, baseTo string) string {
	if decimalValue == 0 {
		return "0"
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
