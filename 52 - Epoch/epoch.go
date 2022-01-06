package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)   // seconds
	fmt.Println(millis) // milliseconds
	fmt.Println(nanos)  // nanoseconds

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
