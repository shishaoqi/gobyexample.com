package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	<-timer.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		for {
			select {
			case <-timer2.C:
				fmt.Println(time.Now())
			default:
				fmt.Println("default: ", time.Now())
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
