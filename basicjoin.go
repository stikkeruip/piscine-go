package piscine

func BasicJoin(elems []string) string {
	s := ""

	for _, i := range elems {
		s = s + i
	}

	return s
}
