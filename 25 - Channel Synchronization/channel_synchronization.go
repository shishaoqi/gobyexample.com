package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	// The done channel will be used to notify another goroutine
	// that this function's work is done.
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	//time.Sleep(time.Second)

	// If you removed the <- done line from this program,
	// the program would exit before the worker even started.
	<-done
}
