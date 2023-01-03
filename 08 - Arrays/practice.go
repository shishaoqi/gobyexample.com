package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("empty array: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get:", a[4])

	fmt.Println("array length:", len(a))

	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare a array:", a2)

	var twoDimension [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimension[i][j] = i + j
		}
	}
	fmt.Println(twoDimension)
}
