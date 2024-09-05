package main

import "fmt"

func BasicAtoi2(s string) int {
	rns := []rune(s)
	result := 0
	for i := 0; i < len(rns); i++ {
		if rns[i] >= '0' && rns[i] <= '9' {
			result = result*10 + int(rns[i]-'0')
		} else {
			result = 0
			break
		}
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
