package main

import "fmt"

func main() {
	a := 5
	var b int
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	ans := addition(a, b)
	fmt.Println("sum is = ", ans)
}

func addition(a, b int) int {
	sum := a + b
	return sum
}
