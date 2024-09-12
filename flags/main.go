package main

import (
	"fmt"
	"os"
)

func main() {
	str := ""
	getOrder := false
	getInsert := false
	strToInsert := ""

	if len(os.Args) == 1 {
		printHelp()
		return
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--order" || os.Args[i] == "-o" {
			getOrder = true
		} else if os.Args[i] == "--insert" || os.Args[i] == "-i" {
			getInsert = true
			strToInsert = os.Args[i+1]
			i++
		} else if len(os.Args[i]) >= 9 && os.Args[i][:9] == "--insert=" {
			getInsert = true
			strToInsert = os.Args[i][9:]
		} else if os.Args[i][:3] == "-i=" && len(os.Args[i]) >= 3 {
			getInsert = true
			strToInsert = os.Args[i][3:]
		} else {
			str = str + os.Args[i]
		}
	}

	if getInsert {
		str = str + strToInsert
	}
	if getOrder {
		str = orderString(str)
	}

	fmt.Print(str)
}

func orderString(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}

	return string(runes)
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	This flag will behave like a boolean, if it is called it will order the argument.")
}
