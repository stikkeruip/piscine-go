package printcombn

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n > 9 {
		return
	}
	comb := make([]rune, n)
	first := true
	var recurse func(int, int)
	recurse = func(index, start int) {
		if index == n {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			for _, digit := range comb {
				z01.PrintRune(digit)
			}
			first = false
			return
		}
		for i := start; i <= 9; i++ {
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
