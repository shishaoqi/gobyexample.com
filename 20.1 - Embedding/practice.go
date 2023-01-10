package main

import "fmt"

type base struct {
	num int
}

func (b base) describe(n int) string {
	return fmt.Sprintf("base with number: %v", n)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "hello",
	}
	fmt.Printf("co={num:%v, str:%v}\n", co.num, co.str)
	fmt.Println("alse num:", co.base.num)
	fmt.Println("describe:", co.describe(8888))

	type describer interface {
		describe(int) string // 函数参数与返回参数都得加上
	}

	var d describer = co
	fmt.Println("describer:", d.describe(666))
}
