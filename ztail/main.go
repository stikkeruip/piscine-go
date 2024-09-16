package main

import (
	"fmt"
	"os"
)

func checkFile(e error, name string) bool {
	if e != nil {
		fmt.Printf("open %s: no such file or directory\n", name)
		return false
	}
	return true
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . -c <number> <file1> <file2> ...")
		os.Exit(1)
	}

	// Parse the number of bytes to read
	numStr := os.Args[2]
	num := 0
	for _, r := range numStr {
		num = num*10 + int(r-'0')
	}

	errExit := false
	firstFile := true
	prevHadError := false

	for _, f := range os.Args[3:] {
		file, err := os.Open(f)
		if checkFile(err, f) {
			// Handle newlines before the file header
			if !firstFile {
				// If the previous output was an error, print one newline
				if prevHadError {
					fmt.Printf("\n")
				} else {
					// If the previous output was not an error, print one newline
					fmt.Printf("\n")
				}
			}
			// Print the header
			fmt.Printf("==> %s <==\n", f)

			// Read the file content
			data, _ := os.ReadFile(f)
			start := len(data) - num
			if start < 0 {
				start = 0
			}
			fmt.Printf("%s", data[start:])

			file.Close()
			prevHadError = false
		} else {
			// An error occurred
			errExit = true
			prevHadError = true
		}
		firstFile = false
	}

	if errExit {
		os.Exit(1)
	}
}
