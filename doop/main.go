package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	strNum1 := os.Args[1]
	op := os.Args[2]
	strNum2 := os.Args[3]

	num1 := Atoi(strNum1)
	num2 := Atoi(strNum2)

	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int {
			if b == 0 {
				fmt.Println("No division by 0")
				os.Exit(0)
			}
			return a / b
		},
		"%": func(a, b int) int {
			if b == 0 {
				fmt.Println("No modulo by 0")
				os.Exit(0)
			}
			return a % b
		},
	}

	if operation, exists := ops[op]; exists {
		// Perform the operation
		result := operation(num1, num2)

		// Convert the result to a string and print each rune
		printInt(result)
	} else {
		return // Return early if the operator is invalid
	}
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	rns := []rune(s)
	sign := 1
	start := 0
	result := 0

	// Detect sign
	if rns[0] == '+' {
		start = 1
	} else if rns[0] == '-' {
		start = 1
		sign = -1
	}

	// Define the maximum and minimum int values (for 32-bit integers)
	const MaxInt32 = 2147483647
	const MinInt32 = -2147483648

	for i := start; i < len(rns); i++ {
		if rns[i] >= '0' && rns[i] <= '9' {
			digit := int(rns[i] - '0')

			// Check for potential overflow before multiplying and adding the digit
			if sign == 1 && (result > (MaxInt32-digit)/10) {
				os.Exit(0) // Return 0 for overflow
			}
			if sign == -1 && (result > (-MinInt32-digit)/10) {
				os.Exit(0) // Return 0 for overflow
			}

			result = result*10 + digit
		} else {
			// Invalid character, return 0 for error
			os.Exit(0)
		}
	}

	return result * sign
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Collect digits
	digits := []rune{}
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}

	// Print digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
	z01.PrintRune('\n')
}
