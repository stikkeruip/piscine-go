package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	first := true
	for a := 0; a < 10; a++ {
		for b := a + 1; b < 10; b++ {
			for c := b + 1; c < 10; c++ {
				if !first {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(rune(a + '0'))
				z01.PrintRune(rune(b + '0'))
				z01.PrintRune(rune(c + '0'))
				first = false
			}
		}
	}
}
