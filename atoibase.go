package piscine

func AtoiBase(n string, base string) int {
	if !IsValid(n, base) {
		return 0
	}

	if base == "0123456789" {
		return Atoi(n)
	}
	if base == "01" {
		return FromBin(Atoi(n))
	}
	if base == "0123456789ABCDEF" {
		return FromHex(n)
	}
	return FromCustom(n, base)
}

func FromCustom(n string, base string) int {
	baseMap := make(map[rune]int)
	for i, char := range base {
		baseMap[char] = i
	}

	result := 0
	baseLen := len(base)
	for _, char := range n {
		result = result*baseLen + baseMap[char]
	}
	return result
}

func FromHex(n string) int {
	hexMap := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15,
	}

	result := 0
	for _, char := range n {
		result = result*16 + hexMap[char]
	}
	return result
}

func FromBin(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + 2*FromBin(n/10)
}

func IsValid(n string, base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}

	for _, char := range n {
		if !seen[char] {
			return false
		}
	}

	return true
}
