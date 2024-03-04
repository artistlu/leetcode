package count_primes01

func countPrimes(n int) int {
	isPrimes := make([]int, n)

	for i := 2; i*i < n; i++ {
		if isPrimes[i] == 0 {
			for j := i * i; j < n; j = j + i {
				isPrimes[j] = 1
			}
		}
	}

	var counts int
	for i := 2; i < n; i++ {
		if isPrimes[i] == 0 {
			counts++
		}
	}

	return counts
}
