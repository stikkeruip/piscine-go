package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	ints := make([]int, max-min)
	count := 0

	for i := min; i < max; i++ {
		ints[count] = i
		count++
	}

	return ints
}
