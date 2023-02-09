package main

import (
	"fmt"
)

func main() {
	//记住，在go语言中，区别一个变量是数组还是切片，就看有没有定义长度
	var arr1 = [5]int{1, 2, 3, 4, 5}

	fmt.Println("arr1:", arr1)
	fmt.Println("arr1 length:", len(arr1))

	slice1 := arr1[:]
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1 length:", len(slice1))
	fmt.Println("slice1 caption:", cap(slice1))

	slice1 = append(slice1, 6)
	fmt.Println("slice1 new caption:", cap(slice1))

	slice1 = append(slice1, 7, 8, 9, 10, 11, 12)
	fmt.Println("slice1 new caption:", cap(slice1))

	fmt.Println(arr1[1:2])
}
