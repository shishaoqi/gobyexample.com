package main

import (
	"fmt"
)

func main() {

	whatTypeIs := func(v interface{}) {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float64")
		case string:
			fmt.Println("string")
		case interface{}:
			fmt.Println("inteface")
		default:
			fmt.Println("unknow type")
		}
	}

	whatTypeIs("abc")

	whatTypeIs(111)
}
