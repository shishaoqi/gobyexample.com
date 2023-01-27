package main

import (
	"fmt"
	"time"
)

func main() {
	// Tick is a convenience wrapper for NewTicker providing access to the ticking channel only.
	tick := time.Tick(500 * time.Millisecond)
	stop := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		stop <- true
	}()

	for {
		select {
		case t := <-tick:
			fmt.Println(t)
		case <-stop:
			fmt.Println("reveive stop signal")
			return
		}
	}
}
