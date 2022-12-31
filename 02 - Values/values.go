package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // Strings, which can be added together with +.

	fmt.Println("1+1 = ", 1+1)         // Integers
	fmt.Println("7.0/3.0 = ", 7.0/3.0) // Floats

	// Booleans, with boolean operators as you'd expect
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
