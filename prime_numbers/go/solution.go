package main

import (
	"fmt"
)

/* Prime numbers */

// Naive
func isPrime(n int) bool {
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

// Sqrt
func isSqrtPrime(n int) bool {
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

func test() {
	fmt.Println(isPrime(0) && isSqrtPrime(0))
	fmt.Println(isPrime(1) && isSqrtPrime(1))
	fmt.Println(isPrime(2) && isSqrtPrime(2))
	fmt.Println(isPrime(3) && isSqrtPrime(3))
	fmt.Println(!isPrime(4) && !isSqrtPrime(4))
}

func main() {
}
