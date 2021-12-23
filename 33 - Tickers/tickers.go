package main

import (
	"fmt"
	"time"
)

// ticker <俚>心脏 -- 时钟ticker.ticker
func main() {
	ticker := time.NewTicker(500 * time.Microsecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(3600 * time.Microsecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
