package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	first := true
	for a := 0; a <= 99; a++ {
		for b := a + 1; b <= 99; b++ {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))
			first = false
		}
	}
}
