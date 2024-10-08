// 2. program to find sum of fibonaci series upto n.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms in the Fibonacci series: ")
	fmt.Scanln(&n)

	sum := fibonacciSum(n)
	fmt.Printf("sum of the first %d terms of the Fibonacci series = %d\n", n, sum)
}

// calculate the sum of Fibonacci series - n terms
func fibonacciSum(n int) int {
	if n <= 0 {
		return 0
	}

	a, b := 0, 1
	sum := 0
	// Calculate Fibonacci numbers and their sum
	for i := 0; i < n; i++ {
		sum += a
		a, b = b, a+b
	}

	return sum
}
