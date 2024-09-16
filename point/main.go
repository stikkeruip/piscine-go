package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Function to print a number as runes
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append(digits, rune(digit)+'0')
		n /= 10
	}

	// Print digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func main() {
	points := &point{}
	setPoint(points)

	strx := "x = "
	stry := ", y = "

	// Print "x = "
	for _, r := range strx {
		z01.PrintRune(r)
	}
	// Print point.x value
	printNumber(points.x)

	// Print ", y = "
	for _, r := range stry {
		z01.PrintRune(r)
	}
	// Print point.y value
	printNumber(points.y)

	// Print a newline
	z01.PrintRune('\n')
}
