package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var ints []int

	for i := min; i < max; i++ {
		ints = append(ints, i)
	}

	return ints
}
