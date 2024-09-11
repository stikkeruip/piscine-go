package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	offsetUpper := 0
	start := 1

	if os.Args[1] == "--upper" {
		offsetUpper = -32
		start = 2
	}

	intToRune := map[int]rune{
		1: 'a', 2: 'b', 3: 'c', 4: 'd',
		5: 'e', 6: 'f', 7: 'g', 8: 'h', 9: 'i',
		10: 'j', 11: 'k', 12: 'l', 13: 'm', 14: 'n',
		15: 'o', 16: 'p', 17: 'q', 18: 'r', 19: 's',
		20: 't', 21: 'u', 22: 'v', 23: 'w', 24: 'x',
		25: 'y', 26: 'z',
	}

	for i := start; i < len(os.Args); i++ {
		arg := os.Args[i]
		num := Atoi(arg)
		if num >= 1 && num <= 26 {
			z01.PrintRune(intToRune[num] + rune(offsetUpper))
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	rns := []rune(s)
	result := 0

	for i := 0; i < len(rns); i++ {
		if rns[i] >= '0' && rns[i] <= '9' {
			result = result*10 + int(rns[i]-'0')
		} else {
			result = 0
			break
		}
	}
	return result
}
