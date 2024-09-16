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
	numStr := os.Args[2]
	num := 0
	errExit := false

	for _, r := range numStr {
		num = num*10 + int(r-'0')
	}

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
			fmt.Printf("\n==> %s <==", f)
			fmt.Printf("\n%s", data[start:])
		} else {
			errExit = true
		}
	}

	if errExit {
		os.Exit(1)
	}
}
