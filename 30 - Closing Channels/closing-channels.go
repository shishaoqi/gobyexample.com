package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			fmt.Println("more: ", j, more)
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 4; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	//close(jobs) // 注释掉后，会有奇效
	fmt.Println("sent all jobs")

	<-done
}
