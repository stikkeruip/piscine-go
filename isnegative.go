package isnegative

import "github.com/01-edu/z01"

func isnegative(nb int) {
	if(nb >= 0) {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
	if(nb < 0) {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	}
}