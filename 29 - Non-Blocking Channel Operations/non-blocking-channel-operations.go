package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("reveive message:", msg)
	default:
		fmt.Println("no message reveived")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}
	// go func() {
	// 	messages <- msg
	// }()
	// time.Sleep(time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
