// Write a program to find prime numbers in a given range

package main

import (
	"fmt"
	"math"
)

func main() {
	var startNumber int
	fmt.Print("Enter the start number: ")
	fmt.Scanln(&startNumber)
	var endNumber int
	fmt.Print("Enter the end number: ")
	fmt.Scanln(&endNumber)
	fmt.Printf("Prime numbers between %d and %d are:\n", startNumber, endNumber)
	for i := startNumber; i <= endNumber; i++ {
		if checkPrime(i) == true {
			fmt.Println(i)
		}
	}
}

func checkPrime(n int) bool {
	if n < 2 {
		return false // 0 and 1 - not prime numbers
	}
	if n == 2 || n == 3 {
		return true // 2 and 3 are prime numbers
	}
	if n%2 == 0 {
		return false // even no are not prime
	}

	for j := 3; j <= int(math.Sqrt(float64(n))); j += 2 { //checking for odd numbers only
		if n%j == 0 {
			return false
		}
	}
	return true
}
