package main

import "fmt"

func Atoi(s string) int {
	rns := []rune(s)
	sign := 1
	start := 0
	result := 0

	if rns[0] == '+' {
		start = 1
	}

	if rns[0] == '-' {
		start = 1
		sign = -1
	}

	for i := start; i < len(rns); i++ {
		if rns[i] >= '0' && rns[i] <= '9' {
			result = result*10 + int(rns[i]-'0')
		} else {
			result = 0
			break
		}
	}
	return result * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
