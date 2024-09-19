package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	first := true
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := i * 3; j < 3+(i*3); j++ {
			if first {
				first = false
				fmt.Printf("%d", deck[j])
				continue
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
			fmt.Printf("%d", deck[j])
		}
		first = true
		z01.PrintRune('\n')
	}
}
