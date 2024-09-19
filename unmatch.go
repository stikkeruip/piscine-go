package piscine

func Unmatch(a []int) int {
	seen := make(map[int]int)

	for _, n := range a {
		seen[n]++
	}

	unmatched := -1
	for key, value := range seen {
		if value%2 != 0 {
			if unmatched == -1 {
				unmatched = key
				continue
			}
			if findAscii(key) < findAscii(unmatched) {
				unmatched = key
			}
		}
	}

	return unmatched
}

func findAscii(num int) int {
	sum := 0

	if num < 0 {
		num = -num
	}

	for num > 0 {
		digit := num % 10        // Get the last digit
		asciiValue := digit + 48 // Convert the digit to its ASCII value
		sum += asciiValue        // Add the ASCII value to the sum
		num /= 10                // Remove the last digit
	}

	return sum
}
