package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // creates a map with a sring key and int value

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // print the whole map

	v1 := m["k1"] // gets th value for the key k1
	fmt.Println("v1:", v1)

	v3 := m["k3"] 
	fmt.Println("v3:", v3) // if key does not exist, a zero value is returned

	fmt.Println("len:", len(m)) // builtin 	len rturns the # of key/val pairs

	delete(m, "k2") // builtin delete removes key/value pairs from a map
	fmt.Println("map:", m)

	clear(m) // removes all key/val pairs
	fmt.Println("m:", m)

	_, prs := m["k2"] // optional second return value when getting a value from a map returns a bool if the key is present or not
	fmt.Println("prs:", prs) // we dont need the value itself so we use a blank identifier _


	n := map[string]int{"foo" : 1, "bar" : 2} // here we initialize and declare a map in the same line
	fmt.Println("map:", n)

	n2 := map[string]int{"foo" : 1, "bar" : 2}
	if maps.Equal(n, n2) { // maps package contins useful util functions. Here we use Equal to check if two maps are the same
		fmt.Println("n == n2")
	}
}
