package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5} // declaring array
	for k, v := range a {
		fmt.Println("array: ", k, " = ", v)
	}

	// Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4} // slice
	sum := 0
	for _, num := range nums { // using _ (blank identifier) to ignore the index
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("maps: %s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range 在字符串中迭代 unicode 码点（code point)。±
	// 第一个返回值是字符(rune)的起始字节位置，然后第二个是字符本身(rune)
	//
	for i, c := range "golang" {
		fmt.Println("index:", i, "character: ", c)
	}
}
