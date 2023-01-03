package main

import "fmt"

// 流程控制(控制结构)中的循环语句：for，另外还有 for - range(遍历语句)
// for is Go's only looping constuct

func main() {
	// go 的 for 循环类型于 C 的，但又不完全相同。它统一了 for 和 while, 并且没有 do-while。
	// 有三种形式，其中只有一个具有分号。
	/**
	1. Like a C for
	---------------------
	for init; condition; post {}
	*/
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	/**
	2. Like a C while
	---------------------
	for condtion {}
	*/
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	/**
	3. Like a C for（;;），
	for(;;){}是产生无限循环的惯用方法，其只能使用 break、goto 或 return 来结束。
	---------------------
	for {}
	*/
	for {
		fmt.Println("loop")
		break
	}

	// continue 运用
	for n := 0; n <= 9; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// 标签的作用对象为外循环，因此 i 会直接变成下一个循环的值，
	// 而此时 j 的值就被重设为 0 。
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d \n", i, j)
		}
	}

	// goto
	i = 0
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
