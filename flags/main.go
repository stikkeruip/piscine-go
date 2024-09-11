package main

import "os"

func main() {
	if os.Args[1] == "--insert" || os.Args[1] == "-i" {
		insertString()
	}
	if os.Args[1] == "--order" || os.Args[1] == "-o" {
		orderString()
	}
}

func insertString() {

}

func orderString() {

}
