package main

import "fmt"

func main() {
	a := 0
	b := &a
	n := &b
	**n = 1
	fmt.Println(a)
}
