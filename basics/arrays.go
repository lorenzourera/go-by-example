package main

import "fmt"

func main() {

	var a [5]int // create an array that holds 5 ints, by default zero-valued
	fmt.Println("emp:", a)

	a[4] = 100 // makes the 5th value (start at 0) = 100
	fmt.Println("set:", a) // shows the entire array
	fmt.Println("get:", a[4]) // shows only the 5th element (n = 4, n+1)

	fmt.Println("len:", len(a)) // len returns the length of an array

	b := [5]int{1, 2, 3, 4, 5} // here we declare & initialize an array of length 5 with numbers 1 to 5
	fmt.Println("dcl", b)

	b = [...]int{1, 2, 3, 4, 5} // we can also make the compiler count the number of elements for us
	fmt.Println("idx:", b)

	var twoD [2][3]int // two dimensional 2x3 array using nested for loops to build it
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD =  [2][3]int{ // you can create and initialize multi-dimensional arrays at once too
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
