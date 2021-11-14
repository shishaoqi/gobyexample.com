package main

import "fmt"

// map 是 Go 内建的关联数据类型（在一些其他语言中也被称为 哈希（hash）或者 字典（dict））
func main() {
	// 要创建一个空 map，需要使用内建函数 make:
	// make(map[key-type]val-type)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	//delete(m, "k2")
	fmt.Println("map:", m)

	val, present := m["k2"]
	fmt.Println("val , present:", val, present)

	// 另一种语法声明并初始化一个新的 map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
