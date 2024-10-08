package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter number: ")
	var n int
	fmt.Scanln(&n)

	if isPrime(n) {
		fmt.Println(n, "is a Prime number")
	} else {
		fmt.Println(n, "is Not a Prime number")
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false // 0 and 1
	}
	if n == 2 || n == 3 {
		return true // 2 and 3 - prime numbers
	}
	if n%2 == 0 {
		return false // Even numbers - not prime
	}

	// Check for odd divisors only up to the square root of n
	for j := 3; j <= int(math.Sqrt(float64(n))); j += 2 {
		if n%j == 0 {
			return false
		}
	}
	return true
}
