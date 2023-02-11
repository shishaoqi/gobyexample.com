package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.23", 32) // the result still has type float64, but it will be convertible to float32 without changing its value.
	fmt.Printf("%T\n", f)
	fmt.Printf("%v, %v\n", f, float32(f))

	f1, _ := strconv.ParseFloat("1.23", 64) // the result still has type float64, but it will be convertible to float32 without changing its value.
	fmt.Printf("%T\n", f1)
	fmt.Printf("%v, %v\n", f1, float32(f1))

	i1, r := strconv.ParseInt("101", 2, 64)
	if r != nil {
		fmt.Println(r)
	}
	fmt.Println(i1)

	i2, r2 := strconv.ParseInt("1c8", 16, 64) // 0x1c8 会报错
	if r2 != nil {
		fmt.Println(r2)
	}
	fmt.Println(i2)

	i3, r3 := strconv.ParseInt("0x1c8", 0, 64)
	if r3 != nil {
		fmt.Println(r3)
	}
	fmt.Println(i3)

	i4, _ := strconv.ParseUint("123", 0, 64)
	fmt.Println(i4)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
