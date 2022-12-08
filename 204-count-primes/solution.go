package countprimes

import "math"

// to understand how to solve this problem, let's first take a journey back to ancient Greece and learn about the Sieve of Eratosthenes
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

func countPrimes(n int) int {
	// anything less than 2 has no primes as defined by the contraints: 0 <= n <= 5 * 106
	if n <= 2 {
		return 0
	}

	// make a list of numbers 0 to n
	potentialPrimes := make([]bool, n)
	for i := 0; i < n; i++ {
		potentialPrimes[i] = true
	}

	potentialPrimes[0] = false // 0 is not prime
	potentialPrimes[1] = false // 1 is not prime

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if prime := potentialPrimes[i]; prime {
			for j := i * i; j < n; j += i {
				potentialPrimes[j] = false
			}
		}
	}

	count := 0
	for _, isPrime := range potentialPrimes {
		if isPrime {
			count++
		}
	}

	return count
}
