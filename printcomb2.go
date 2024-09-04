package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	first := true
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					if !first {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(a + '0'))
					z01.PrintRune(rune(b + '0'))
					z01.PrintRune(' ')
					z01.PrintRune(rune(c + '0'))
					z01.PrintRune(rune(d + '0'))
					first = false
				}
			}
		}
	}
}
