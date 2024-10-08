// Problem Statement:
// Create  a greeting app, which greets the user depending on what time it is.

// Good morning! : When the time is between 6:00:01 am to 11:00:00 am
// Good afternoon! : When the time is between 11:00:01 am to 4:00:00 pm
// Good evening! : When the time is between 4:00:01 pm to 9:00:00 pm
// Good night! : When the time is between 9:00:01 pm to 6:00:00 am

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	hour := currentTime.Hour()

	// Create a greeting based on the time
	var greeting string
	switch {
	case hour >= 6 && hour < 11:
		greeting = "Good morning!"
	case hour >= 11 && hour < 16:
		greeting = "Good afternoon!"
	case hour >= 16 && hour < 21:
		greeting = "Good evening!"
	default:
		greeting = "Good night!"
	}

	fmt.Println(greeting)
}
