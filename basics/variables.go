package main

import "fmt"

func main(){
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // you specify that you are declaring a variable, then you specify the type, then the value
	fmt.Println(b, c)
	
	var d = true
	fmt.Println(d)

	var e int // variables declared but not initialized will be zero-valued. for an int the zero value is 
	fmt.Println(e)

	f := "apple" // infers the variable type based on the value
	fmt.Println(f)
	g := "7j"
	fmt.Println(g)

}