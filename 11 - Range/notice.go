package main

import (
	"fmt"
)

func main() {
	var s = []int{1, 2, 3, 4}
	var m = make(map[int]*int)

	/** -- 经典考题
	新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[k] = &v 取的都是变量 v 的地址。
	map 中的所有元素的值都是变量 v 的地址，故所有 v 被赋值为 0xc0000aa008，所有输出都是 0xc0000aa008.
	*/
	for k, v := range s {
		fmt.Println(k, " ===> ", v)
		fmt.Println(k, " ===> ", &v)

		//tmp := v
		//m[k] = &tmp
		m[k] = &v
	}

	for key, val := range m {
		fmt.Println(key, " ===> ", val)
	}
}
