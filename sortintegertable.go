package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
