package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	file := getBaseFileName(os.Args[0])

	for _, i := range file {
		z01.PrintRune(i)
	}
}

func getBaseFileName(path string) string {
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
