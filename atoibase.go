package piscine

func AtoiBase(n string, base string) int {
	if base == "0123456789" {
		return Atoi(n)
	}
	return 0
}
