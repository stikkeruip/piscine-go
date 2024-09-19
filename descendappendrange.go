package piscine

func DescendAppendRange(max, min int) []int {
	if max < min {
		return []int{}
	}

	var numSlice []int
	if max == min {
		numSlice = append(numSlice, max)
	}

	for i := max; i > min; i-- {
		numSlice = append(numSlice, i)
	}

	return numSlice
}
