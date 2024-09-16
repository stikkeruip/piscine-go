package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("Wrong file name")
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	check(err)
	fmt.Print(string(data))
}
