package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	for _, s := range os.Args {
		if s == "01" || s == "galaxy" || s == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
