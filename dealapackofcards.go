package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	first := true
	people := 6
	evenCards := len(deck) / people

	dealCards(deck, people, evenCards, first)
}

func dealCards(deck []int, people int, evenCards int, first bool) {
	for i := 0; i < people; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := i * evenCards; j < evenCards+(i*evenCards); j++ {
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
