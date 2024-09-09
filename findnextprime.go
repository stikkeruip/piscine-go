package piscine

func FindNextPrime(n int) int {
	if n <= 2 {
		return 2
	}

	num := n

	if n%2 == 0 {
		num++
	}

	for {
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return num
		}
		num += 2
	}
}
