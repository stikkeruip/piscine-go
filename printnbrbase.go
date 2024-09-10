package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = nbr * -1
	}
	if base == "0123456789" {
		printDigits(nbr)
		return
	}
	if base == "0123456789ABCDEF" {
		ToHex(nbr)
		return
	}
	if base == "01" {
		ToBin(nbr)
		return
	}
	ToCustom(nbr, base)
	return
}

func ToHex(n int) {
	if n == 0 {
		return
	}
	ToHex(n / 16)
	digit := n % 16
	if digit < 10 {
		z01.PrintRune(rune(digit + '0'))
	} else {
		z01.PrintRune(rune(digit - 10 + 'A'))
	}
}

func ToBin(n int) {
	if n == 0 {
		return
	}

	ToBin(n / 2)
	digit := n % 2
	z01.PrintRune(rune(digit + '0'))
}

func ToCustom(n int, s string) {
	if n == 0 {
		return
	}

	ToCustom(n/len(s), s)
	digit := n % len(s)
	z01.PrintRune(rune(s[digit]))
}

func printDigits(n int) {
	if n > 0 {
		if n/10 > 0 {
			printDigits(n / 10)
		}
		z01.PrintRune(rune(n%10 + '0'))
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	if base[0] == '-' || base[0] == '+' {
		return false
	}

	seen := make(map[rune]bool)
	uniqueCount := 0

	for _, char := range base {
		if !seen[char] {
			seen[char] = true
			uniqueCount++
		}
	}

	return uniqueCount >= 2
}
