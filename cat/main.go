package main

import (
	"bufio"
	"fmt"
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
	// filenames := os.Args[1:]
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		for _, r := range message {
			z01.PrintRune(r)
		}
		return
	}
	for _, f := range os.Args[1:] {
		data, err := os.ReadFile(f)
		if checkFile(err, f) {
			fmt.Println(data)
		}
	}
}
