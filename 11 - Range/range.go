package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for k, v := range a {
		fmt.Println(k, " = ", v)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
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
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range 在字符串中迭代 unicode 码点（code point)。±
	// 第一个返回值是字符的起始字节位置，然后第二个是字符本身
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
