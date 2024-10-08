// 3. program to find count of even, odd and zero

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scanln(&n)

	evenCount := 0
	oddCount := 0
	zeroCount := 0

	fmt.Println("Enter the numbers:")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)

		// Count even, odd, and zero
		if num == 0 {
			zeroCount++
		} else if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Printf("Count of Even numbers: %d\n", evenCount)
	fmt.Printf("Count of Odd numbers: %d\n", oddCount)
	fmt.Printf("Count of Zeroes: %d\n", zeroCount)
}
