package main

import "fmt"

func main() {

	// if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if only
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if only + or statement
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// else if, else
	if num := 9; num < 9 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}