package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	sLen := len(s)
	toFindLen := len(toFind)

	if toFindLen > sLen {
		return -1
	}

	for i := 0; i <= sLen-toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		}
	}

	return -1
}
