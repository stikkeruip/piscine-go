package piscine

func Compact(ptr *[]string) int {
	count := 0
	compacted := []string{}
	for _, s := range *ptr {
		if s != "" {
			compacted = append(compacted, s)
			count++
		}
	}
	*ptr = compacted
	return count
}
