package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var str []string

	for i := 1; i < len(os.Args); i++ {
		str = append(str, os.Args[i])
	}

	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if str[j] > str[j+1] {
				str[j], str[j+1] = str[j+1], str[j]
			}
		}
	}

	for _, i := range str {
		for _, j := range i {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
