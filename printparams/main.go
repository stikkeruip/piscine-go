package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		for _, i := range os.Args[i] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
