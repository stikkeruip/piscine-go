package main

import "fmt"

func BasicAtoi(s string) int {
	rns := []rune(s)
	result := 0
	for i := 0; i < len(rns); i++ {
		result = result*10 + int(rns[i]-'0')
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
