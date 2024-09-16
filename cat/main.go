package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func checkFile(e error, name string) bool {
	if e != nil {
		str := "ERROR: open " + name + ": No such file or directory"
		for _, r := range str {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
		return false
	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		input, _ := io.ReadAll(os.Stdin)

		for _, r := range string(input) {
			z01.PrintRune(r)
		}
		return
	}
	for _, f := range os.Args[1:] {
		data, err := os.ReadFile(f)
		if checkFile(err, f) {
			for _, r := range string(data) {
				z01.PrintRune(r)
			}
		}
	}
	os.Exit(1)
}
