package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var runes []rune

	for i := 1; i < len(os.Args); i++ {
		runes = append(runes, []rune(os.Args[i])[0])
	}

	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}

	for _, i := range runes {
		z01.PrintRune(i)
	}
}
