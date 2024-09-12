package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}

	str := ""
	for i := 1; i < len(os.Args); i++ {
		str += os.Args[i] + " "
	}

	runes := []rune(str)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !containsVowel(runes[left]) {
			left++
		}
		for left < right && !containsVowel(runes[right]) {
			right--
		}
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	for i := 0; i < len(runes)-1; i++ {
		z01.PrintRune(runes[i])
	}
}

func containsVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}
