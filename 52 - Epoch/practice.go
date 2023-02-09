package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())

	p(time.Unix(0, now.Unix()))
	p(time.Unix(now.UnixNano(), 0))
}
