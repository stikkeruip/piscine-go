package main

import (
	"fmt"
	"os"
)

func checkFile(e error, name string) bool {
	if e != nil {
		fmt.Printf("open %s: no such file or directory", name)
		return false
	}
	return true
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . -c <number> <file1> <file2> ...")
		os.Exit(1)
	}

	numStr := os.Args[2]
	num := 0
	errExit := false

	for _, r := range numStr {
		num = num*10 + int(r-'0')
	}

	firstFile := true

	for _, f := range os.Args[3:] {
		file, err := os.Open(f)
		if checkFile(err, f) {
			fileInfo, _ := file.Stat()
			size := fileInfo.Size()
			start := size - int64(num)
			if start < 0 {
				start = 0
			}
			data, _ := os.ReadFile(f)

			// Print a newline between files, except before the first file
			if !firstFile {
				fmt.Printf("\n")
			}

			// Print the header and content with appropriate newlines
			fmt.Printf("==> %s <==\n", f)
			fmt.Printf("%s", data[start:])

			firstFile = false
		} else {
			// Handle file open error
			if !firstFile {
				fmt.Printf("\n")
			}
			errExit = true
			firstFile = false
		}
	}

	if errExit {
		os.Exit(1)
	}
}
