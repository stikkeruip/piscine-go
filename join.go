package piscine

func Join(strs []string, sep string) string {
	s := ""

	for v, i := range strs {
		s = s + i

		if v == len(strs)-1 {
			return s
		}

		s = s + sep
	}

	return s
}
