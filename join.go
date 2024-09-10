package piscine

func Join(strs []string, sep string) string {
	s := ""

	for _, i := range strs {
		s = s + i + sep
	}

	return s
}
