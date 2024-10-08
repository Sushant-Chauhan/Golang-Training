// Write a program to find the second largest number in an array.

package main

import "fmt"

func main() {
	firstLargest := -9999999
	secondLargest := -9999999
	arr := [5]int{50, 8, 18, 28, 38}
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] > firstLargest {
			secondLargest = firstLargest
			firstLargest = arr[i]
		} else if arr[i] > secondLargest && firstLargest > arr[i] {
			secondLargest = arr[i]
		}
	}

	fmt.Println("Second largest number is = ", secondLargest)
}
