package piscine

import "github.com/01-edu/z01"

const size = 8

func printSolution(queens []int) {
	for i := 0; i < size; i++ {
		z01.PrintRune(rune(queens[i] + '1'))
	}
	z01.PrintRune('\n')
}

func isSafe(queens []int, currRow, currCol int) bool {
	for row := 0; row < currRow; row++ {
		col := queens[row]
		if col == currCol || currRow-row == currCol-col || currRow-row == col-currCol {
			return false
		}
	}
	return true
}

func solve(queens []int, currentRow int) {
	if currentRow == size {
		printSolution(queens)
		return
	}

	for col := 0; col < size; col++ {
		if isSafe(queens, currentRow, col) {
			queens[currentRow] = col
			solve(queens, currentRow+1)
		}
	}
}

func EightQueens() {
	queens := make([]int, size)
	solve(queens, 0)
}
