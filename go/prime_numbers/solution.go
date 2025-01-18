// Algorithm: Search for prime numbers
package algo

// Method: Naive
func IsPrimeNumber(n int) bool {
	if n < 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Method: SQRT
func IsPrimeNumberWithSqrt(n int) bool {
	if n < 2 {
		return true
	}

	for i := 2; (i * i) <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
