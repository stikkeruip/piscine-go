package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a > 0; a-- {
		for b := a - 1; b >= 0; b-- {
			if a == 01 && b == 00 {
				z01.PrintRune(rune(a/10 + '0'))
				z01.PrintRune(rune(a%10 + '0'))
				break
			}
			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))

			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
