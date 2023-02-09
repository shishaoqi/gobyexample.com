package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	// Printf 格式化  -- %02d：数字位数不足 2 位时，左边补零。
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// Sprintf 格式化
	timeStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(timeStr)

	// now.Format 格式化
	p(now.Format("2006-01-02 15:04:05"))
	p(now.Format("2006/01/02 15:04:05"))
	p(now.Format("2006-01-02")) // 年月日
	p(now.Format("15:04:05"))   // 时分秒
}
