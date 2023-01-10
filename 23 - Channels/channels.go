package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines

func main() {
	// Create a new channel with make(chan val-type)
	messages := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
	// By default sends and reveives block until both the sender and
	// reveiver are ready
}
