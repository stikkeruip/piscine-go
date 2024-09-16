package main

import (
	"fmt"
	"os"
)

func checkFile(e error, name string) {
	if e != nil {
		fmt.Printf("ERROR: open %s: No such file or directory", name)
	}
}

func main() {
	// filenames := os.Args[1:]

	for _, f := range os.Args[1:] {
		data, err := os.ReadFile(f)
		checkFile(err, f)
		fmt.Println(data)
	}
}
