package main

import (
	"fmt"
	"time"
)

// 好好体会下此处：fatal error: all goroutines are asleep - deadlock! 问题的解决

func worker(jobs, results chan int) {
	for j := range jobs {
		fmt.Printf("worker %v start job\n", j)
		time.Sleep(time.Second)
		results <- j * 100
		fmt.Printf("worker %v finsh job\n", j)
	}
}

func main() {
	const jobNum = 5

	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)

	for i := 0; i <= 3; i++ {
		go worker(jobs, results)
	}

	for j := 1; j <= 5; j++ { // 把等号去掉，会报错，为什么呢
		jobs <- j
	}
	close(jobs)

	// 输出方案一：不确认有多少数量，一直接收
	// for r := range results {
	// 	fmt.Printf("result of results: %v\n", r)
	// }

	// 输出方案二：把数量确定好
	// for a := 1; a <= jobNum; a++ {
	// 	<-results
	// }

	// 输出方案三：一直接收（也不确认有多少），但会在接收到最后一个后把 channel 关闭
	for {
		j, more := <-results
		fmt.Println("more: ", j, more)
		if more {
			fmt.Println("receive job", j)
			if j == 500 {
				close(results)
			}
		} else {
			fmt.Println("receive all jobs")
			//close(results)
			return
		}
	}
}
