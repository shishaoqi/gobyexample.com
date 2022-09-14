package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // int 可省略
	fmt.Println(b, c)

	var d = true // 省略了标识类型符 bool
	fmt.Println(d)

	var e int // 没有初始变量的值时，则必须在变量名后加上类型，此时变量会默认加上该类型的初始值
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a varible
	f := "short"
	fmt.Println(f)
}