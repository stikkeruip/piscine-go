package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n > 9 {
		return
	}
	comb := make([]rune, n)
	var recurse func(int, int)
	recurse = func(index, start int) {
		if index == n {
			for _, digit := range comb {
				z01.PrintRune(digit)
			}
			if comb[0] != rune('0'+10-n) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}
		for i := start; i <= 10-n+index; i++ {
			comb[index] = rune(i + '0')
			recurse(index+1, i+1)
		}
	}
	recurse(0, 0)
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(3)
}

