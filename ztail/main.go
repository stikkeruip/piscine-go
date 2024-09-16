package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func checkFile(e error, name string) bool {
	if e != nil {
		str := "open " + name + ": no such file or directory"
		for _, r := range str {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
		return false
	}
	return true
}

func main() {
	numStr := os.Args[2]
	num := 0

	for _, r := range numStr {
		num = num*10 + int(r-'0')
	}

	for _, f := range os.Args[3:] {
		file, err := os.Open(f)
		if checkFile(err, f) {
			fileInfo, _ := file.Stat()
			size := fileInfo.Size()
			data, _ := os.ReadFile(f)
			fmt.Printf("==> %s <==\n", f)
			fmt.Printf("%s\n", data[size-int64(num):])
		} else {
			os.Exit(1)
		}
	}
}
