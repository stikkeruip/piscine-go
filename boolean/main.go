package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) int {
	if nbr%2 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	lengthOfArg := len(os.Args)

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
