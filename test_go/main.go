package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	var temp1 int = *a / *b
	var temp2 int = *a % *b

	*a = temp1
	*b = temp2
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
