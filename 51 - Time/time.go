package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Year())
	p(now.Month())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Nanosecond()) // 纳秒偏移
	p(now.Location())
	p(now.Weekday())

	p("--------------------")
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	p("--------------------")
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
