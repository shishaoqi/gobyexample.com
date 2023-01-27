package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker *time.Ticker
	ticker = time.NewTicker(500 * time.Microsecond)
	done := make(chan string)

	go func() {
		for {
			select {
			case <-done:
			case t := <-ticker.C:
				fmt.Println("time at:", t)
			}
		}
	}()

	time.Sleep(5000 * time.Microsecond)
	done <- "done"
	ticker.Stop()
	fmt.Println("ticker stop")
}
