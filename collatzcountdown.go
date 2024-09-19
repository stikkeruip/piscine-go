package piscine

func CollatzCountdown(start int) int {
	count := 0

	for start != 1 {
		if start%2 == 0 {
			start /= 2
			count++
		} else {
			start = (start * 3) + 1
			count++
		}
	}
	return count
}
