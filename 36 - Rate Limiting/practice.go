package main

import (
	"fmt"
	"time"
)

func main() {
	// fisrt limiter -- 生产在前，限制消费在后
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)
	ticker := time.Tick(500 * time.Millisecond)
	for x := range request {
		<-ticker
		//fmt.Println(x)
		fmt.Println("request", x, time.Now())
	}

	// another limiter -- 限制消费在前，生产在后
	fmt.Println("another limiter ---")
	l2 := make(chan time.Time, 2) // 最小不能小于下面的最小值
	for i := 0; i < 2; i++ {
		l2 <- time.Now()
	}
	go func() {
		for t := range time.Tick(800 * time.Millisecond) {
			l2 <- t
		}
	}()
	//close(l2) // 关闭后，有神奇效果

	r2 := make(chan int, 9)
	for i := 0; i < 9; i++ {
		r2 <- i
	}
	close(r2)
	for r := range r2 {
		<-l2
		fmt.Println("request", r, time.Now())
	}
}
