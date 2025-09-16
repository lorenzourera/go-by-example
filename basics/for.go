package main

import "fmt"

func main() {

	// Basic Type - single condition
	fmt.Println("Single Condition For Loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic - initial/condition/after for loop
	fmt.Println("Classic")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// range over integer, or do this N times
	fmt.Println("Range Over Integer")
	for i := range 3 {
		fmt.Println("range", i)
	}

	// No condition - will iterate until you breakout of the loop or return a variable
	fmt.Println("No Condition")
	for {
		fmt.Println("loop")
		break
	}

	// Use continue
	fmt.Println("Use continue")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}