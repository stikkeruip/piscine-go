package piscine

func UltimateDivMod(a *int, b *int) {
	var temp1 int = *a / *b
	var temp2 int = *a % *b

	*a = temp1
	*b = temp2
}
