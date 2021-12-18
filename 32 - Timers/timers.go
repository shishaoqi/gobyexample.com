package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(2 * time.Second)

	<-timer.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	//stop2 := timer2.Stop()
	//if stop2 {
	//	fmt.Println("timer 2 stopped")
	//}

	time.Sleep(4 * time.Second)
}
